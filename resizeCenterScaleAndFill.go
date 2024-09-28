package main

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/nfnt/resize"
)

// sampleBackgroundColor samples a pixel from the top-left corner as the background color
func sampleBackgroundColor(img image.Image) color.Color {
	// Sample from the top-left corner or any other desired point
	return img.At(0, 0)
}

// resizeCenterScaleAndFill resizes, centers, scales down the content, and fills the background with a sampled color.
func resizeCenterScaleAndFill(img image.Image, targetWidth, targetHeight int) image.Image {
	// Step 1: Resize and center the content
	centeredImg := resizeAndCenter(img, targetWidth, targetHeight)

	// Step 2: Scale down by 50% (2x smaller)
	scaledWidth := targetWidth / 2
	scaledHeight := targetHeight / 2

	// Resize the centered image to 50% of its size
	scaledImg := resize.Resize(uint(scaledWidth), uint(scaledHeight), centeredImg, resize.Lanczos3)

	// Step 3: Sample a background color from the original image
	backgroundColor := sampleBackgroundColor(img)

	// Step 4: Create the final image and fill the background with the sampled color
	finalImg := image.NewRGBA(image.Rect(0, 0, targetWidth, targetHeight))
	draw.Draw(finalImg, finalImg.Bounds(), &image.Uniform{backgroundColor}, image.Point{}, draw.Src)

	// Step 5: Place the scaled image back in the center of the final image
	offsetX := (targetWidth - scaledWidth) / 2
	offsetY := (targetHeight - scaledHeight) / 2
	draw.Draw(finalImg, image.Rect(offsetX, offsetY, offsetX+scaledWidth, offsetY+scaledHeight), scaledImg, image.Point{}, draw.Over)

	return finalImg
}
