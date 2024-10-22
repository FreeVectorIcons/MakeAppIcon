package main

import (
	"fmt"
	"image"
	"path/filepath"
)

func generateAssets(platforms *Platforms) {
	for _, platform := range platforms.Platforms {
		fmt.Printf("Generating assets for platform: %s\n", platform.Title)

		for _, bundle := range platform.BundleSpecs {
			fmt.Printf(" - Generating %s\n", bundle.Category)

			if platform.ID == "android" {
				generateAndroidAssets(platform, bundle, platforms.SourceIcon, platforms.SourceSplash)
				continue
			}

			sourceImageFile := platforms.SourceIcon
			if bundle.Category == "splash" {
				sourceImageFile = platforms.SourceSplash
			}

			sourceImage, err := openImage(sourceImageFile)
			if err != nil {
				fmt.Printf("Error opening source image: %v\n", err)
				continue
			}

			for _, file := range bundle.ImageSet {
				finalWidth := file.Width * file.Scale
				finalHeight := file.Height * file.Scale
				var resizedImage image.Image

				if bundle.ResizeFromCenter {
					// Use the resizeCenterScaleAndFill strategy
					resizedImage = resizeCenterScaleAndFill(sourceImage, finalWidth, finalHeight)
				} else {
					// Use default resizing strategy
					resizedImage = resizeAndCenter(sourceImage, finalWidth, finalHeight)
				}

				err := saveGeneratedAsset(platform.Path, bundle.Path, file.Filename, resizedImage)
				if err != nil {
					fmt.Printf("Error saving image: %v\n", err)
				}
			}

			if platform.ID == "ios" && bundle.Category == "icon" {
				manifestPath := filepath.Join(platform.Path, bundle.Path, "Contents.json")
				err := generateManifest(bundle.ImageSet, manifestPath)
				if err != nil {
					fmt.Printf("Error generating manifest: %v\n", err)
				}
			}
		}
	}
}
