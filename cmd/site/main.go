package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	log.SetFlags(0)
	if len(os.Args) != 2 || (os.Args[1] != "build" && os.Args[1] != "dev") {
		log.Fatal("usage: go run ./cmd/site [build|dev]")
	}
	if err := build("www", "templates", "dist"); err != nil {
		log.Fatal(err)
	}
	if os.Args[1] == "dev" {
		if err := develop(); err != nil {
			log.Fatal(err)
		}
	}
}

func build(source, partials, output string) error {
	shared, err := template.ParseGlob(filepath.Join(partials, "*.html"))
	if err != nil {
		return err
	}
	files := map[string][]byte{}
	err = filepath.WalkDir(source, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}
		if !entry.Type().IsRegular() {
			return fmt.Errorf("unsupported source file: %s", path)
		}
		relative, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".html") {
			page, err := shared.Clone()
			if err != nil {
				return err
			}
			page, err = page.New("page").Parse(string(data))
			if err != nil {
				return fmt.Errorf("%s: %w", path, err)
			}
			var rendered bytes.Buffer
			if err := page.Execute(&rendered, struct{ Path string }{filepath.ToSlash(relative)}); err != nil {
				return fmt.Errorf("%s: %w", path, err)
			}
			data = rendered.Bytes()
		}
		files[relative] = data
		return nil
	})
	if err != nil {
		return err
	}
	// Render every page before changing the output so template errors preserve the last build.
	for relative, data := range files {
		target := filepath.Join(output, relative)
		existing, err := os.ReadFile(target)
		if err == nil && bytes.Equal(existing, data) {
			continue
		}
		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			return err
		}
		if err := os.WriteFile(target, data, 0644); err != nil {
			return err
		}
	}
	return filepath.WalkDir(output, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}
		relative, err := filepath.Rel(output, path)
		if err != nil {
			return err
		}
		if _, exists := files[relative]; !exists {
			return os.Remove(path)
		}
		return nil
	})
}

func fingerprint() ([32]byte, error) {
	hash := sha256.New()
	for _, root := range []string{"templates", "www"} {
		err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if entry.IsDir() {
				return nil
			}
			info, err := entry.Info()
			if err != nil {
				return err
			}
			fmt.Fprintf(hash, "%s\x00%d\x00%d\n", path, info.Size(), info.ModTime().UnixNano())
			return nil
		})
		if err != nil {
			return [32]byte{}, err
		}
	}
	var result [32]byte
	copy(result[:], hash.Sum(nil))
	return result, nil
}

func develop() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	previous, err := fingerprint()
	if err != nil {
		return err
	}
	command := exec.CommandContext(ctx, "go", "tool", "devd", "-l", "./dist")
	command.Stdin, command.Stdout, command.Stderr = os.Stdin, os.Stdout, os.Stderr
	command.WaitDelay = 3 * time.Second
	if err := command.Start(); err != nil {
		return err
	}
	done := make(chan error, 1)
	go func() { done <- command.Wait() }()
	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case err := <-done:
			if ctx.Err() != nil {
				return nil
			}
			return err
		case <-ticker.C:
			current, err := fingerprint()
			if err != nil {
				log.Printf("watch failed: %v", err)
				continue
			}
			if current == previous {
				continue
			}
			previous = current
			if err := build("www", "templates", "dist"); err != nil {
				log.Printf("build failed: %v", err)
			} else {
				log.Print("site rebuilt")
			}
		}
	}
}
