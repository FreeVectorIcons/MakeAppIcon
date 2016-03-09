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
		var app_icons AppIconContents
		err = json.Unmarshal([]byte(APP_ICON_JSON), &app_icons)
		if err != nil {
			log.Fatal(err)
		}

		// Go thorugh the list of images
		for i := 0; i < len(app_icons.Images); i++ {
			image_info := app_icons.Images[i]

			// Parse scalar size
			size_x, _ := strconv.ParseFloat(strings.Split(image_info.Size, "x")[0], 64)
			scale, _ := strconv.ParseFloat(strings.Split(image_info.Scale, "x")[0], 64)

			app_icons.Images[i].image = resize.Resize(uint(size_x*scale), 0, img, resize.Lanczos3)
		}

		// Save
		app_icons.Save(".")
	}

	app.Run(os.Args)
}
