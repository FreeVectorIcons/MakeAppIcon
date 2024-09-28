package main

import (
	"fmt"
	"image"
	"path/filepath"
)

func generateAndroidAssets(platform PlatformSpec, bundle BundleSpec) {
	sourceImageFile := platform.SourceIcon
	if bundle.Category == "splash" {
		sourceImageFile = platform.SourceSplash
	}

	sourceImage, err := openImage(sourceImageFile)
	if err != nil {
		fmt.Printf("Error opening source image: %v\n", err)
		return
	}

	for _, file := range bundle.ImageSet {
		// Adjust width and height based on scale
		finalWidth := file.Width * file.Scale
		finalHeight := file.Height * file.Scale
		var resizedImage image.Image

		// Check if resizeFromCentre is true in the spec
		if bundle.ResizeFromCenter {
			// Use the resizeCenterScaleAndFill strategy
			resizedImage = resizeCenterScaleAndFill(sourceImage, finalWidth, finalHeight)
		} else {
			// Use default resizing strategy
			resizedImage = resizeAndCenter(sourceImage, finalWidth, finalHeight)
		}

		// Determine the output folder
		var folder string
		if file.Density != "" {
			folder = filepath.Join(platform.Path, file.Density) // drawable-xxhdpi, mipmap-hdpi, etc.
		} else {
			folder = filepath.Join(platform.Path, "drawable") // default drawable directory
		}

		// Save the generated image
		err := saveGeneratedAsset(folder, "", file.Filename, resizedImage)
		if err != nil {
			fmt.Printf("Error saving image: %v\n", err)
		}
	}
}
