package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

// Helper function to create a simple image for testing
func createTestImage(width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.RGBA{uint8(x % 256), uint8(y % 256), 0, 255})
		}
	}
	return img
}

func TestConvertToRGBA(t *testing.T) {
	img := createTestImage(100, 100)
	rgbaImg := convertToRGBA(img)

	if rgbaImg.Bounds() != img.Bounds() {
		t.Errorf("Expected bounds %v, but got %v", img.Bounds(), rgbaImg.Bounds())
	}
}

func TestResizeAndCenter(t *testing.T) {
	img := createTestImage(200, 200)
	resized := resizeAndCenter(img, 100, 100)

	if resized.Bounds().Dx() != 100 || resized.Bounds().Dy() != 100 {
		t.Errorf("Expected dimensions 100x100, but got %dx%d", resized.Bounds().Dx(), resized.Bounds().Dy())
	}

	// Save the result for manual inspection (optional)
	file, err := os.Create("resized_test_output.png")
	if err != nil {
		t.Fatalf("Failed to create output file: %v", err)
	}
	defer file.Close()
	png.Encode(file, resized)
}
