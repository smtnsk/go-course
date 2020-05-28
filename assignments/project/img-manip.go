// TODO:
//	greyscale
//	shrink / expand to given dimensions or multiplier
//	manipulate color channels independently
//	blur
//	edge detection
//	spawn a goroutine per pixel

package main

import (
    "os"
	"math"
	"image"
	"image/jpeg"
	"image/png"
	"image/color"
    "path/filepath"
	"fmt"
	"log"
	"strings"
)

func check_error(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
	// Check parameters
	if len(os.Args) != 2 {
		log.Fatalln("ERROR: give a single file path as an argument")
	}
	img_path := os.Args[1] // os.Args[0] == program filename

	// Open file
	file, err := os.Open(img_path)
	check_error(err)

	// Decode file
	img, format, err := image.Decode(file)
	check_error(err)

	// Create a new image that matches the original's dimensions
	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	new_img := image.NewRGBA(rect)

	// Loop vertically through original image
	for y := 0; y < size.Y; y++ {
		// Loop horizontally through original image
		for x := 0; x < size.X; x++ {
			// Get pixel from original image at current coordinates
			pixel := img.At(x, y)

			// Get the pixel's original color values
			orig_color := color.RGBAModel.Convert(pixel).(color.RGBA)

			// Redden the image
			r := uint8(math.Max(0, math.Min(float64(orig_color.R) * 1.5, 255)))
			g := uint8(math.Max(0, math.Min(float64(orig_color.G) * 0.5, 255)))
			b := uint8(math.Max(0, math.Min(float64(orig_color.B) * 0.5, 255)))
			a := orig_color.A

			// Convert into proper color data
			c := color.RGBA{ R: r, G: g, B: b, A: a, }

			// Set the color of the corresponding pixel in the new image
			new_img.Set(x, y, c)
		}
	}
	file.Close()

	// Create filename (<filename>-modified.<ext>)
	ext := filepath.Ext(img_path)
	name := strings.TrimSuffix(filepath.Base(img_path), ext)
	path := filepath.Dir(img_path)
	new_img_path := fmt.Sprintf("%s/%s-modified%s", path, name, ext)

	// Create a new file
	new_file, err := os.Create(new_img_path)
	check_error(err)

	// Save the new image to the new file
	if format == "jpeg"{
		check_error(jpeg.Encode(new_file, new_img, nil))
	} else {
		check_error(png.Encode(new_file, new_img))
	}
	check_error(new_file.Close())
}
