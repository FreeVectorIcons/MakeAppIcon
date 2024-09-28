package main

import (
	"image"
	"image/draw"

	"github.com/nfnt/resize"
)

// convertToRGBA converts any image.Image to *image.RGBA
func convertToRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, src, image.Point{}, draw.Src)
	return rgba
}

// resizeAndCenter resizes the image to fit within the target dimensions and centers the content, cropping any excess.
func resizeAndCenter(img image.Image, targetWidth, targetHeight int) image.Image {
	originalWidth := img.Bounds().Dx()
	originalHeight := img.Bounds().Dy()

	// Calculate aspect ratios
	originalAspectRatio := float64(originalWidth) / float64(originalHeight)
	targetAspectRatio := float64(targetWidth) / float64(targetHeight)

	var newWidth, newHeight int
	if originalAspectRatio > targetAspectRatio {
		// Image is wider than target, scale by height and crop width
		newHeight = targetHeight
		newWidth = int(float64(targetHeight) * originalAspectRatio)
	} else {
		// Image is taller than target, scale by width and crop height
		newWidth = targetWidth
		newHeight = int(float64(targetWidth) / originalAspectRatio)
	}

	// Resize the image to the new dimensions while maintaining aspect ratio
	resizedImg := resize.Resize(uint(newWidth), uint(newHeight), img, resize.Lanczos3)

	// Convert the resized image to RGBA (if necessary)
	resizedRGBA := convertToRGBA(resizedImg)

	// Calculate the coordinates for cropping to center the content
	cropX := (newWidth - targetWidth) / 2
	cropY := (newHeight - targetHeight) / 2

	// Crop the resized image to fit the target dimensions exactly
	croppedImg := resizedRGBA.SubImage(image.Rect(cropX, cropY, cropX+targetWidth, cropY+targetHeight)).(*image.RGBA)

	return croppedImg
}
