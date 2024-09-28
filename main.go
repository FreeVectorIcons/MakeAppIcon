// ./MakeAppIcon -filename assets/appIcon1024.png -splash assets/splash4096.png
package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/nfnt/resize"
)

// App Icon structs
type AppIconItem struct {
	Size     string `json:"size"`
	Idiom    string `json:"idiom"`
	Filename string `json:"filename"`
	Scale    string `json:"scale"`

	image image.Image
}

type AppIconContents struct {
	Images []AppIconItem `json:"images"`
	Info   struct {
		Version int    `json:"version"`
		Author  string `json:"author"`
	}
}

// AndroidIcon and WindowsIcon structs
type AndroidIcon struct {
	Density  string
	Filename string
	image    image.Image
}

type WindowsIcon struct {
	Filename string
	image    image.Image
}

// Load the iOS App Icon template from JSON
func loadAppIconTemplate() AppIconContents {
	templateFile := "appicon_template.json"
	data, err := ioutil.ReadFile(templateFile)
	if err != nil {
		log.Fatal("Error reading the app icon template file:", err)
	}

	var appIconContents AppIconContents
	err = json.Unmarshal(data, &appIconContents)
	if err != nil {
		log.Fatal("Error unmarshalling the app icon template JSON:", err)
	}

	return appIconContents
}

// Resize iOS app icons
func resizeAppIcons(img image.Image, icons []AppIconItem) []AppIconItem {
	for i, icon := range icons {
		parts := strings.Split(icon.Size, "x")
		width, _ := strconv.Atoi(parts[0])
		height, _ := strconv.Atoi(parts[1])

		resizedImage := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
		icons[i].image = resizedImage
	}
	return icons
}

// Save iOS icons
func (icon AppIconContents) Save(dir string) {
	sep := string(os.PathSeparator)
	path := fmt.Sprintf("%s%sImages.xcassets%sAppIcon.appiconset", dir, sep, sep)

	err := os.MkdirAll(path, 0777)
	if err != nil {
		log.Fatal(err)
	}

	for _, imageInfo := range icon.Images {
		if imageInfo.image != nil {
			newPath := fmt.Sprintf("%s%s%s", path, sep, imageInfo.Filename)
			out, err := os.Create(newPath)
			if err != nil {
				log.Print(err)
			}
			defer out.Close()

			png.Encode(out, imageInfo.image)
		}
	}

	contentPath := fmt.Sprintf("%s%sContents.json", path, sep)
	contentFile, err := os.Create(contentPath)
	if err != nil {
		log.Fatal(err)
	}
	defer contentFile.Close()

	encoder := json.NewEncoder(contentFile)
	encoder.SetIndent("", "    ")
	encoder.Encode(icon)
}

// Generate Android icons
func generateAndroidIcons(img image.Image) []AndroidIcon {
	densities := []string{"ldpi", "mdpi", "hdpi", "xhdpi", "xxhdpi", "xxxhdpi"}
	sizes := map[string]int{"ldpi": 36, "mdpi": 48, "hdpi": 72, "xhdpi": 96, "xxhdpi": 144, "xxxhdpi": 192}
	var icons []AndroidIcon

	for _, density := range densities {
		size := sizes[density]
		resizedImage := resize.Resize(uint(size), uint(size), img, resize.Lanczos3)

		// For both drawable and mipmap directories
		icons = append(icons, AndroidIcon{Density: fmt.Sprintf("drawable-%s", density), Filename: "icon.png", image: resizedImage})
		icons = append(icons, AndroidIcon{Density: fmt.Sprintf("mipmap-%s", density), Filename: "icon.png", image: resizedImage})
	}
	return icons
}

// Generate Android splash screens
func generateAndroidSplashScreens(img image.Image) []AndroidIcon {
	densities := []string{"ldpi", "mdpi", "hdpi", "xhdpi", "xxhdpi", "xxxhdpi"}
	sizes := map[string][2]int{
		"ldpi":    {320, 240},
		"mdpi":    {480, 320},
		"hdpi":    {800, 480},
		"xhdpi":   {1280, 720},
		"xxhdpi":  {1600, 960},
		"xxxhdpi": {1920, 1280},
	}
	var splashScreens []AndroidIcon

	for _, density := range densities {
		width, height := sizes[density][0], sizes[density][1]
		resizedImage := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

		// For both drawable and drawable-land directories
		splashScreens = append(splashScreens, AndroidIcon{Density: fmt.Sprintf("drawable-%s", density), Filename: "splashscreen_image.png", image: resizedImage})
		splashScreens = append(splashScreens, AndroidIcon{Density: fmt.Sprintf("drawable-land-%s", density), Filename: "splashscreen_image.png", image: resizedImage})
	}
	return splashScreens
}

// Generate Windows icons
func generateWindowsIcons(img image.Image) []WindowsIcon {
	iconSizes := []int{16, 24, 30, 32, 42, 48, 50, 54, 70, 90, 120, 150, 210, 256, 270}
	var icons []WindowsIcon

	for _, size := range iconSizes {
		resizedImage := resize.Resize(uint(size), uint(size), img, resize.Lanczos3)
		icons = append(icons, WindowsIcon{Filename: fmt.Sprintf("iconApplicationIcon-%dx%d.png", size, size), image: resizedImage})
	}
	return icons
}

// Generate Windows splash screens
func generateWindowsSplashScreens(img image.Image) []WindowsIcon {
	splashScreenSizes := []struct {
		width  int
		height int
	}{
		{620, 300},
		{868, 420},
		{1116, 540},
	}
	var splashScreens []WindowsIcon

	for _, size := range splashScreenSizes {
		resizedImage := resize.Resize(uint(size.width), uint(size.height), img, resize.Lanczos3)
		splashScreens = append(splashScreens, WindowsIcon{Filename: fmt.Sprintf("screenSplashscreen-%dx%d.png", size.width, size.height), image: resizedImage})
	}
	return splashScreens
}

// Main function
func main() {
	app := cli.NewApp()
	app.Name = "MakeAppIcon"
	app.Version = "0.6"
	app.Usage = "CLI tool to make app icons and splash screens for iOS, Android, and Windows"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "filename",
			Value: "Ex: icon.png",
			Usage: "PNG icon file of size 1024x1024",
		},
		cli.StringFlag{
			Name:  "splash",
			Value: "Ex: splashscreen.png",
			Usage: "PNG splash screen file of size 4096x4096",
		},
	}

	app.Action = func(c *cli.Context) {
		iconFile := c.String("filename")
		splashFile := c.String("splash")

		// Open and validate app icon file
		icon, err := os.Open(iconFile)
		if err != nil {
			log.Fatal(err)
		}
		imgIcon, err := png.Decode(icon)
		if err != nil {
			log.Fatal(err)
		}
		icon.Close()

		// Validate the icon size
		iconSize := imgIcon.Bounds()
		if !(iconSize.Max.X == 1024 && iconSize.Max.Y == 1024) {
			log.Fatal("App icon must be 1024x1024.")
		}

		// Open and validate splash screen file
		splash, err := os.Open(splashFile)
		if err != nil {
			log.Fatal(err)
		}
		imgSplash, err := png.Decode(splash)
		if err != nil {
			log.Fatal(err)
		}
		splash.Close()

		// Validate the splash screen size
		splashSize := imgSplash.Bounds()
		if !(splashSize.Max.X == 4096 && splashSize.Max.Y == 4096) {
			log.Fatal("Splash screen must be 4096x4096.")
		}

		// Load JSON templates for iOS, Android, and Windows
		appIconContents := loadAppIconTemplate()
		androidIcons := generateAndroidIcons(imgIcon)
		androidSplashScreens := generateAndroidSplashScreens(imgSplash)
		windowsIcons := generateWindowsIcons(imgIcon)
		windowsSplashScreens := generateWindowsSplashScreens(imgSplash)

		// Generate iOS assets
		appIconContents.Images = resizeAppIcons(imgIcon, appIconContents.Images)
		appIconContents.Save("output/iOS")

		// Generate Android assets (icons and splash screens)
		SaveAndroidIcons("output/android", androidIcons)
		SaveAndroidIcons("output/android", androidSplashScreens)

		// Generate Windows assets (icons and splash screens)
		SaveWindowsAssets("output/windows", windowsIcons, windowsSplashScreens)
	}

	app.Run(os.Args)
}

// Save Android icons and splash screens
func SaveAndroidIcons(dir string, icons []AndroidIcon) {
	sep := string(os.PathSeparator)
	basePath := fmt.Sprintf("%s%sres", dir, sep)

	for _, icon := range icons {
		path := fmt.Sprintf("%s%s%s%s%s", basePath, sep, icon.Density, sep, icon.Filename)
		os.MkdirAll(fmt.Sprintf("%s%s%s", basePath, sep, icon.Density), 0777)
		out, err := os.Create(path)
		if err != nil {
			log.Print(err)
		}
		defer out.Close()

		png.Encode(out, icon.image)
	}
}

// Save Windows icons and splash screens
func SaveWindowsAssets(dir string, icons []WindowsIcon, splashScreens []WindowsIcon) {
	sep := string(os.PathSeparator)
	iconPath := fmt.Sprintf("%s%sicons", dir, sep)
	splashscreenPath := fmt.Sprintf("%s%ssplashscreens", dir, sep)

	// Make directories for icons and splash screens
	os.MkdirAll(iconPath, 0777)
	os.MkdirAll(splashscreenPath, 0777)

	// Save icons
	for _, icon := range icons {
		path := fmt.Sprintf("%s%s%s", iconPath, sep, icon.Filename)
		out, err := os.Create(path)
		if err != nil {
			log.Print(err)
		}
		defer out.Close()

		png.Encode(out, icon.image)
	}

	// Save splash screens
	for _, splash := range splashScreens {
		path := fmt.Sprintf("%s%s%s", splashscreenPath, sep, splash.Filename)
		out, err := os.Create(path)
		if err != nil {
			log.Print(err)
		}
		defer out.Close()

		png.Encode(out, splash.image)
	}
}
