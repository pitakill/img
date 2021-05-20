package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
)

func main() {
	f := flag.String("f", "", "File name of the image")
	flag.Parse()

	if *f == "" {
		fmt.Println("Please provide an image with the flag `-f`")
		return
	}

	path, err := filepath.Abs(*f)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Assume the image is jpg
	// Dirty and Nasty, sorry :-P
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}

	bounds := img.Bounds()
	i := image.NewRGBA(bounds)
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			i.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
		}
	}

	filename := *f

	extension := filepath.Ext(*f)
	name := filename[0 : len(filename)-len(extension)]

	output, err := os.Create(name + "_gray" + extension)
	if err != nil {
		log.Fatalln(err)
	}
	defer output.Close()

	jpeg.Encode(output, i, nil)
}
