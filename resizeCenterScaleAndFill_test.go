package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

// Test sampleBackgroundColor to ensure it samples the correct color
func TestSampleBackgroundColor(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	sampleColor := color.RGBA{255, 0, 0, 255} // Red color
	img.Set(0, 0, sampleColor)

	resultColor := sampleBackgroundColor(img)
	if resultColor != sampleColor {
		t.Errorf("Expected color %v, but got %v", sampleColor, resultColor)
	}
}

// Test resizeCenterScaleAndFill for proper scaling and filling
func TestResizeCenterScaleAndFill(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 200, 200))
	targetWidth, targetHeight := 100, 100

	// Fill with red to observe background filling
	red := color.RGBA{255, 0, 0, 255}
	for y := 0; y < 200; y++ {
		for x := 0; x < 200; x++ {
			img.Set(x, y, red)
		}
	}

	resized := resizeCenterScaleAndFill(img, targetWidth, targetHeight)

	if resized.Bounds().Dx() != targetWidth || resized.Bounds().Dy() != targetHeight {
		t.Errorf("Expected dimensions %dx%d, but got %dx%d",
			targetWidth, targetHeight, resized.Bounds().Dx(), resized.Bounds().Dy())
	}

	// Save the result for manual inspection (optional)
	file, err := os.Create("resized_scaled_filled_test_output.png")
	if err != nil {
		t.Fatalf("Failed to create output file: %v", err)
	}
	defer file.Close()
	png.Encode(file, resized)
}
