package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

type AppIconImage struct {
	Size     string `json:"size"`
	Idiom    string `json:"idiom"`
	Filename string `json:"filename"`
	Scale    string `json:"scale"`

	image image.Image
}

type AppIconContents struct {
	Images []AppIconImage `json:"images"`
	Info   struct {
		Version int    `json:"version"`
		Author  string `json:"author"`
	}
}

func (icon AppIconContents) Save(dir string) {
	// Create a path
	sep := string(os.PathSeparator)
	path := fmt.Sprintf("%s%sImages.xcassets%sAppIcon.appiconset", dir, sep, sep)

	// Make a folder called Images.xcassets
	err := os.MkdirAll(path, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// Copy all the images
	for _, image_info := range icon.Images {
		if image_info.image != nil {
			new_path := fmt.Sprintf("%s%s%s", path, sep, image_info.Filename)
			out, err := os.Create(new_path)
			if err != nil {
				log.Print(err)
			}
			defer out.Close()
			png.Encode(out, image_info.image)
		}
	}

	// Copy contents.json
	contents_file_path := fmt.Sprintf("%s%sContents.json", path, sep)
	data, err := json.Marshal(icon)
	ioutil.WriteFile(contents_file_path, data, 0644)
}
