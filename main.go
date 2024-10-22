package main

import (
	_ "embed"
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Embed the default spec.yaml
//
//go:embed spec.yaml
var defaultSpec []byte

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
	ID          string       `yaml:"id"`
	Path        string       `yaml:"path"`
	Title       string       `yaml:"title"`
	BundleSpecs []BundleSpec `yaml:"bundleSpecs"`
}

// Platforms wraps the list of platform specifications
type Platforms struct {
	SourceIcon   string         `yaml:"sourceIcon"`
	SourceSplash string         `yaml:"sourceSplash"`
	Platforms    []PlatformSpec `yaml:"platforms"`
}

// saveGeneratedAsset saves the resized image to the specified path
func saveGeneratedAsset(basePath, folder, filename string, img image.Image) error {
	fullPath := filepath.Join(basePath, folder)
	if err := os.MkdirAll(fullPath, 0777); err != nil {
		return fmt.Errorf("error creating directories: %v", err)
	}

	filePath := filepath.Join(fullPath, filename)
	outFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer outFile.Close()

	if err := png.Encode(outFile, img); err != nil {
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

// loadPlatformSpecs parses the YAML data into the Platforms struct
func loadPlatformSpecs(data []byte) (*Platforms, error) {
	var platformSpecs Platforms
	if err := yaml.Unmarshal(data, &platformSpecs); err != nil {
		return nil, fmt.Errorf("error unmarshalling YAML: %v", err)
	}
	return &platformSpecs, nil
}

// updateAssets updates the YAML spec with user-provided icon and splash paths
func updateAssets(platforms *Platforms, iconPath, splashPath string) {
	if iconPath != "" {
		platforms.SourceIcon = iconPath
	}
	if splashPath != "" {
		platforms.SourceSplash = splashPath
	}
}

// Check if required assets exist
func checkAssetsDirectory() error {
	requiredFiles := []string{
		"assets/appIcon1024.png",
		"assets/splash4096.png",
	}
	for _, file := range requiredFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return fmt.Errorf("missing required asset: %s", file)
		}
	}
	return nil
}

func main() {
	// Define CLI flags for spec file and custom assets
	specFile := flag.String("spec", "", "Path to a custom spec.yaml file")
	sourceIcon := flag.String("sourceIcon", "", "Path to a custom source icon")
	sourceSplash := flag.String("sourceSplash", "", "Path to a custom splash image")
	flag.Parse()

	// Load the spec.yaml
	var specData []byte
	var err error

	if *specFile != "" {
		specData, err = os.ReadFile(*specFile)
		if err != nil {
			log.Fatalf("Error reading spec file: %v", err)
		}
		fmt.Println("Using custom spec file:", *specFile)
	} else {
		specData = defaultSpec
		fmt.Println("Using embedded default spec.yaml")
	}

	platforms, err := loadPlatformSpecs(specData)
	if err != nil {
		log.Fatalf("Error loading platform specs: %v", err)
	}

	iconPath := platforms.SourceIcon
	splashPath := platforms.SourceSplash

	if *sourceIcon != "" {
		iconPath = *sourceIcon
		fmt.Println("Using custom source icon:", iconPath)
	}

	if *sourceSplash != "" {
		splashPath = *sourceSplash
		fmt.Println("Using custom splash image:", splashPath)
	}

	if *sourceIcon != "" && *sourceSplash != "" {
		// copy to assets directory
		makeDir("assets")
		copyFile(*sourceIcon, "assets/appIcon1024.png")
		copyFile(*sourceSplash, "assets/splash4096.png")
		fmt.Println("Copied source icon and splash to assets directory")
	}

	if err := checkAssetsDirectory(); err != nil {
		fmt.Printf(`Missing required assets. Please create the "assets" directory with the following structure:
assets/
├── appIcon1024.png
└── splash4096.png
%v`, err)
		panic(err)
	}

	// Update the spec with custom icon and splash paths if provided
	updateAssets(platforms, iconPath, splashPath)

	// Generate assets based on the updated spec
	generateAssets(platforms)
}
