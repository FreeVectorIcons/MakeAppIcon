package main

import (
	"encoding/json"
	"github.com/codegangsta/cli"
	"github.com/nfnt/resize"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
)

// Flags
var filename string
var shouldInstall bool

func main() {
	app := cli.NewApp()
	app.Name = "MakeAppIcon"
	app.Version = "0.1"
	app.Usage = "CLI tool to make app icons for IOS and Android"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "filename",
			Value: "Ex: icon.png",
			Usage: "PNG icon file of size 1024x1024",
		},
	}

	app.Action = func(c *cli.Context) {
		name := ""
		if c.NArg() > 0 {
			name = c.Args()[0]
		}

		//open file
		file, err := os.Open(name)
		if err != nil {
			log.Fatal(err)
		}

		// Decode PNG
		img, err := png.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		// Reject file if bounds is not 1024x1024
		size := img.Bounds()
		if !(size.Max.X == 1024 && size.Max.Y == 1024) {
			log.Fatal("iTunesConnect requires app icon to be of size 1024x1024.")
		}

		// Decode json from the template
		var appIcons AppIconContents
		err = json.Unmarshal([]byte(APP_ICON_JSON), &appIcons)
		if err != nil {
			log.Fatal(err)
		}

		// Go thorugh the list of images
		for i := 0; i < len(appIcons.Images); i++ {
			imageInfo := appIcons.Images[i]

			// Parse scalar size
			sizeX, _ := strconv.ParseFloat(strings.Split(imageInfo.Size, "x")[0], 64)
			scale, _ := strconv.ParseFloat(strings.Split(imageInfo.Scale, "x")[0], 64)

			appIcons.Images[i].image = resize.Resize(uint(sizeX*scale), 0, img, resize.Lanczos3)
		}

		// Save
		appIcons.Save(".")
	}

	app.Run(os.Args)
}
