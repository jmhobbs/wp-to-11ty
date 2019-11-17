package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

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
