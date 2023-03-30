package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

//go:embed all:static
var staticFiles embed.FS

func writeStaticFiles(base string, export *BlogExport) error {
	return fs.WalkDir(staticFiles, "static", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		content, err := staticFiles.ReadFile(path)
		if err != nil {
			return err
		}

		return writeFile(base, strings.TrimPrefix(path, "static/"), content)
	})
}

func writeFile(base, filename string, content []byte) error {
	fullPath := filepath.Join(base, filename)

	if err := os.MkdirAll(filepath.Dir(fullPath), 0700); err != nil {
		return err
	}

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(content)
	return err
}

func writeGlobalDataFiles(base string, export *BlogExport) error {
	dataDir := filepath.Join(base, "_data")
	if err := os.MkdirAll(dataDir, 0700); err != nil {
		return err
	}
	if err := writeBlogData(dataDir, export); err != nil {
		return err
	}
	if err := writeAuthorData(dataDir, export); err != nil {
		return err
	}
	return nil
}

func writeBlogData(dataDir string, export *BlogExport) error {
	f, err := os.Create(filepath.Join(dataDir, "blog.json"))
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}{
		export.Channel.Title,
		export.Channel.Description,
	})
}

func writeAuthorData(dataDir string, export *BlogExport) error {
	f, err := os.Create(filepath.Join(dataDir, "authors.json"))
	if err != nil {
		return err
	}
	defer f.Close()

	authors := make(map[string]Author)
	for _, author := range export.Channel.Authors {
		authors[author.Login] = author
	}

	return json.NewEncoder(f).Encode(authors)
}
