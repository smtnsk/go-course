package main

import (
	"os"
	"image"
	"image/jpeg"
	"image/png"
	"image/color"
	"path/filepath"
	"fmt"
	"strings"
)

func check_error(err error) {
	if err != nil {
		panic(err)
	}
}

func greyscale(img *image.RGBA, width int, height int) {
	// Loop through the canvas in row-major order
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Get the RGB values of the pixel at the current coordinates
			pixel := img.At(x, y)
			r := pixel.(color.RGBA).R
			g := pixel.(color.RGBA).G
			b := pixel.(color.RGBA).B

			// Perform simple greyscaling by averaging the colors
			grey := uint8((float64(r) + float64(g) + float64(b)) / 3.0)

			// Create color struct
			c := color.RGBA{ grey, grey, grey, 255 }

			// Place the colors into the canvas at the same coordinates
			img.Set(x, y, c)
		}
	}
}

func main() {
	// Check parameter(s)
	if len(os.Args) != 2 {
		fmt.Println("ERROR: give a single file path as an argument")
		os.Exit(1)
	}
	img_path := os.Args[1] // os.Args[0] == program filename

	// Open the file
	file, err := os.Open(img_path)
	check_error(err)

	// Decode the file
	img, format, err := image.Decode(file)
	check_error(err)
	file.Close()

	// Get the dimensions of the image
	width := img.Bounds().Size().X
	height := img.Bounds().Size().Y

	//fmt.Printf("File: %s\nFormat: %s\nDimensions: %dx%d\n", img_path, format, width, height)

	// Create a new image (canvas)
	rect := image.Rect(0, 0, width, height)
	new_img := image.NewRGBA(rect)

	// Copy all pixels from the original image onto the canvas;
	// Convert the color space to RGBA
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := img.At(x, y)
			color_rgba := color.RGBAModel.Convert(pixel).(color.RGBA)
			new_img.Set(x, y, color_rgba)
		}
	}

	// Operate on the canvas
	greyscale(new_img, width, height)

	// Create a filename (<filename>-modified.<ext>)
	ext := filepath.Ext(img_path)
	name := strings.TrimSuffix(filepath.Base(img_path), ext)
	path := filepath.Dir(img_path)
	new_img_path := fmt.Sprintf("%s/%s-modified%s", path, name, ext)

	// Create a new file
	new_file, err := os.Create(new_img_path)
	check_error(err)

	// Save the canvas into the new file
	if format == "jpeg"{
		check_error(jpeg.Encode(new_file, new_img, nil))
	} else {
		check_error(png.Encode(new_file, new_img))
	}
	check_error(new_file.Close())
}
