package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ImageEntry struct {
	Idiom    string `json:"idiom"`
	Size     string `json:"size"`
	Scale    string `json:"scale"`
	Filename string `json:"filename"`
}

type InfoEntry struct {
	Version int    `json:"version"`
	Author  string `json:"author"`
}

type Manifest struct {
	Images []ImageEntry `json:"images"`
	Info   InfoEntry    `json:"info"`
}

// generateManifest generates the iOS manifest (Contents.json)
func generateManifest(imageSet []FileSpec, manifestPath string) error {
	var images []ImageEntry

	for _, file := range imageSet {
		sizeStr := fmt.Sprintf("%dx%d", file.Width, file.Height)
		scaleStr := fmt.Sprintf("%dx", file.Scale)
		images = append(images, ImageEntry{
			Idiom:    file.Idiom,
			Size:     sizeStr,
			Scale:    scaleStr,
			Filename: file.Filename,
		})
	}

	manifest := Manifest{
		Images: images,
		Info: InfoEntry{
			Version: 1,
			Author:  "reactnativepro.com",
		},
	}

	// Ensure the directory exists
	err := os.MkdirAll(filepath.Dir(manifestPath), 0777)
	if err != nil {
		return fmt.Errorf("error creating manifest directory: %v", err)
	}

	// Save manifest as JSON
	manifestFile, err := os.Create(manifestPath)
	if err != nil {
		return fmt.Errorf("error creating manifest file: %v", err)
	}
	defer manifestFile.Close()

	encoder := json.NewEncoder(manifestFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(manifest); err != nil {
		return fmt.Errorf("error encoding manifest JSON: %v", err)
	}

	fmt.Printf("Manifest saved to: %s\n", manifestPath)
	return nil
}
