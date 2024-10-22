# MakeAppIcon
A command-line tool to generate app icons for iOS, Android, & windows with ease. This tool simplifies the process of resizing, scaling, and packaging icons required by multiple platforms.

### Install ###

```go get -u github.com/FreeVectorIcons/MakeAppIcon```

### Usage ###

**Usage 1 - With the assets folder**

1.	Create an assets folder in the directory where you’ll run the command.
2.	Copy your icon and splash source images into the assets folder.
3.	Run the following command: ```makeappicon```

**Usage 2 - With source icons directly**

Specify both the app icon and splash image through command-line arguments: `makeappicon -sourceIcon appIcon1024.png -sourceSplash splash4096.png`

**Usage 3 - With a custom YAML specification**

Use a spec.yaml file for advanced customization: `makeappicon -spec spec.yml` See the [spec.yaml](./spec.yaml) reference for an example implementation.


### Output ###

Generates the following files:

```
├── android
│   ├── drawable
│   │   ├── icon.png
│   │   └── splashscreen_image.png
│   ├── drawable-hdpi
│   │   ├── icon.png
│   │   └── splashscreen_image.png
│   ├── drawable-land
│   │   └── splashscreen_image.png
│   ├── drawable-land-hdpi
│   │   └── splashscreen_image.png
│   ├── drawable-land-ldpi
│   │   └── splashscreen_image.png
│   ├── drawable-land-mdpi
│   │   └── splashscreen_image.png
│   ├── drawable-land-xhdpi
│   │   └── splashscreen_image.png
│   ├── drawable-land-xxhdpi
│   │   └── splashscreen_image.png
│   ├── drawable-land-xxxhdpi
│   │   └── splashscreen_image.png
│   ├── drawable-ldpi
│   │   ├── icon.png
│   │   └── splashscreen_image.png
│   ├── drawable-mdpi
│   │   ├── icon.png
│   │   └── splashscreen_image.png
│   ├── drawable-xhdpi
│   │   ├── icon.png
│   │   └── splashscreen_image.png
│   ├── drawable-xxhdpi
│   │   ├── icon.png
│   │   └── splashscreen_image.png
│   ├── drawable-xxxhdpi
│   │   ├── icon.png
│   │   └── splashscreen_image.png
│   ├── mipmap-hdpi
│   │   └── icon.png
│   ├── mipmap-ldpi
│   │   └── icon.png
│   ├── mipmap-mdpi
│   │   └── icon.png
│   ├── mipmap-xhdpi
│   │   └── icon.png
│   ├── mipmap-xxhdpi
│   │   └── icon.png
│   └── mipmap-xxxhdpi
│       └── icon.png
├── iOS
│   └── AppIcon.appiconset
│       ├── Contents.json
│       ├── Icon-20x20@1x.png
│       ├── Icon-20x20@2x.png
│       ├── Icon-20x20@3x.png
│       ├── Icon-29x29@1x.png
│       ├── Icon-29x29@2x.png
│       ├── Icon-29x29@3x.png
│       ├── Icon-40x40@1x.png
│       ├── Icon-40x40@2x.png
│       ├── Icon-40x40@3x.png
│       ├── Icon-60x60@2x.png
│       ├── Icon-60x60@3x.png
│       ├── Icon-76x76@1x.png
│       ├── Icon-76x76@2x.png
│       ├── Icon-83.5x83.5@2x.png
│       └── Icon-marketing-1024x1024.png
└── windows
    ├── icons
    │   ├── iconApplicationIcon-120x120.png
    │   ├── iconApplicationIcon-150x150.png
    │   ├── iconApplicationIcon-16x16.png
    │   ├── iconApplicationIcon-210x210.png
    │   ├── iconApplicationIcon-24x24.png
    │   ├── iconApplicationIcon-256x256.png
    │   ├── iconApplicationIcon-270x270.png
    │   ├── iconApplicationIcon-30x30.png
    │   ├── iconApplicationIcon-32x32.png
    │   ├── iconApplicationIcon-42x42.png
    │   ├── iconApplicationIcon-48x48.png
    │   ├── iconApplicationIcon-50x50.png
    │   ├── iconApplicationIcon-54x54.png
    │   ├── iconApplicationIcon-70x70.png
    │   └── iconApplicationIcon-90x90.png
    └── splashscreens
        ├── screenSplashscreen-1116x540.png
        ├── screenSplashscreen-620x300.png
        └── screenSplashscreen-868x420.png
```

### License ###
The MIT License (MIT)

Copyright (c) 2024 FreeVectorIcons.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

