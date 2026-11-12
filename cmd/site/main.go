package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	log.SetFlags(0)
	if len(os.Args) != 2 || os.Args[1] != "build" {
		log.Fatal("usage: go run ./cmd/site build")
	}
	if err := build("www", "templates", "dist"); err != nil {
		log.Fatal(err)
	}
}

func build(source, partials, output string) error {
	shared, err := template.New("site").Funcs(template.FuncMap{
		"breadcrumbs": breadcrumbTrail,
	}).ParseGlob(filepath.Join(partials, "*.html"))
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

type breadcrumb struct {
	Label string
	URL   string
}

func breadcrumbTrail(parts ...string) ([]breadcrumb, error) {
	if len(parts)%2 != 0 {
		return nil, fmt.Errorf("breadcrumbs require label and URL pairs")
	}
	trail := []breadcrumb{{Label: "Home", URL: "/"}}
	for i := 0; i < len(parts); i += 2 {
		trail = append(trail, breadcrumb{Label: parts[i], URL: parts[i+1]})
	}
	return trail, nil
}
