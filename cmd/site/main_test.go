package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuildAndTemplateRecovery(t *testing.T) {
	root := t.TempDir()
	source, partials, output := filepath.Join(root, "www"), filepath.Join(root, "templates"), filepath.Join(root, "dist")
	write := func(path, content string) {
		t.Helper()
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}
	read := func(path string) string {
		t.Helper()
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		return string(data)
	}
	page := filepath.Join(source, "staff", "index.html")
	header := filepath.Join(partials, "header.html")
	target := filepath.Join(output, "staff", "index.html")
	write(page, `{{template "header.html" .}}<main>Staff</main>`)
	write(header, `<header>{{.Path}}</header>`)
	write(filepath.Join(source, "styles.css"), "body { color: black; }")
	if err := build(source, partials, output); err != nil {
		t.Fatal(err)
	}
	original := read(target)
	if original != "<header>staff/index.html</header><main>Staff</main>" {
		t.Fatal(original)
	}
	if read(filepath.Join(output, "styles.css")) != "body { color: black; }" {
		t.Fatal("asset changed")
	}
	write(header, `{{broken`)
	if err := build(source, partials, output); err == nil {
		t.Fatal("invalid template succeeded")
	}
	if read(target) != original {
		t.Fatal("failed build changed output")
	}
	write(header, `<header>Updated</header>`)
	if err := os.Remove(filepath.Join(source, "styles.css")); err != nil {
		t.Fatal(err)
	}
	if err := build(source, partials, output); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(read(target), "Updated") {
		t.Fatal("partial change was not rendered")
	}
	if _, err := os.Stat(filepath.Join(output, "styles.css")); !os.IsNotExist(err) {
		t.Fatal("stale asset remains")
	}
}

func TestSiteTemplates(t *testing.T) {
	output := t.TempDir()
	if err := build("../../www", "../../templates", output); err != nil {
		t.Fatal(err)
	}
	for _, route := range []string{"index.html", "staff/index.html", "contact/index.html", "festival-artists/index.html", "404.html"} {
		data, err := os.ReadFile(filepath.Join(output, route))
		if err != nil {
			t.Fatal(err)
		}
		html := string(data)
		if strings.Count(html, `<header class="site-header">`) != 1 || strings.Count(html, `<footer class="site-footer">`) != 1 {
			t.Fatalf("missing or duplicate shared content: %s", route)
		}
		if strings.Contains(html, "{{") {
			t.Fatalf("unrendered template: %s", route)
		}
		expected := 1
		if route == "404.html" {
			expected = 0
		}
		if strings.Count(html, `<a aria-current="page"`) != expected {
			t.Fatalf("incorrect active navigation: %s", route)
		}
	}
}
