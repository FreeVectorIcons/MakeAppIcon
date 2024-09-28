package main

import (
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
	"gopkg.in/yaml.v2"
)

// FileSpec defines the individual icon/splash sizes and properties
type FileSpec struct {
	Filename string `yaml:"filename"`
	Width    int    `yaml:"width"`
	Height   int    `yaml:"height"`
	Idiom    string `yaml:"idiom,omitempty"`
	Scale    int    `yaml:"scale,omitempty"`
	Density  string `yaml:"density,omitempty"`
}

// BundleSpec defines a set of files (either icons or splash screens) for a platform
type BundleSpec struct {
	Category            string     `yaml:"category"`
	Path                string     `yaml:"path"`
	Prefix              string     `yaml:"prefix"`
	RemoveAlpha         bool       `yaml:"removeAlpha"`
	MaintainAspectRatio bool       `yaml:"maintainAspectRatio"`
	ResizeFromCenter    bool       `yaml:"resizeFromCenter"`
	ImageSet            []FileSpec `yaml:"imageSet"`
}

// PlatformSpec defines the platform and its corresponding icon and splash files
type PlatformSpec struct {
	ID           string       `yaml:"id"`
	Path         string       `yaml:"path"`
	Title        string       `yaml:"title"`
	SourceIcon   string       `yaml:"sourceIcon"`
	SourceSplash string       `yaml:"sourceSplash"`
	BundleSpecs  []BundleSpec `yaml:"bundleSpecs"`
}

// Platforms wraps the list of platform specifications
type Platforms struct {
	Platforms []PlatformSpec `yaml:"platforms"`
}

// resizeImage resizes the input image to the specified width and height
func resizeImage(img image.Image, width, height int) image.Image {
	return resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
}

// saveGeneratedAsset saves the resized image to the specified path
func saveGeneratedAsset(basePath, folder, filename string, img image.Image) error {
	fullPath := filepath.Join(basePath, folder)
	err := os.MkdirAll(fullPath, 0777)
	if err != nil {
		return fmt.Errorf("error creating directories: %v", err)
	}

	filePath := filepath.Join(fullPath, filename)
	outFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, img)
	if err != nil {
		return fmt.Errorf("error saving PNG: %v", err)
	}

	fmt.Printf("Image saved to: %s\n", filePath)
	return nil
}

// openImage opens a PNG image file and returns an image.Image object
func openImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening image file: %v", err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("error decoding PNG: %v", err)
	}

	return img, nil
}

// loadPlatformSpecs reads the YAML file and parses it into the Platforms struct
func loadPlatformSpecs(filename string) (*Platforms, error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file: %v", err)
	}

	var platformSpecs Platforms
	err = yaml.Unmarshal(yamlFile, &platformSpecs)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling YAML: %v", err)
	}

	return &platformSpecs, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <path-to-yaml-spec>", os.Args[0])
	}

	specFile := os.Args[1]
	platforms, err := loadPlatformSpecs(specFile)
	if err != nil {
		log.Fatalf("Error loading platform specs: %v", err)
	}

	generateAssets(platforms)
}
