package main

import (
 	"fmt"
	"image"
	"image/png"
	"image/color"
	"image/draw"
	"io"
	"os"
	"github.com/oriolad/raytracer/vectors"
)

func main(){
	fmt.Printf("Hello, world.\n")
	var rectangle = image.Rect(0,0,500,500)
	var img = image.NewRGBA(rectangle)
	var uniform = image.NewUniform(color.Black)

	draw.Draw(img, rectangle, uniform, image.ZP, draw.Src)
	//Save image to disk
	fo, err := os.Create("output.png")
	if(err != nil) {
		panic(err)
	}

	var cameraPoint = vectors.Point{0,0,0}
 
	for x := -250.0; x < 250.0; x += 0.1 {
		for y := -250.0; y < 250.0; y += 0.1 {
			var point = vectors.Point{x, y, 1.0}
	        	var ray = vectors.NewLine(cameraPoint, point)	
			var color = trace(ray)

			img.Set(x, y, color)
		}
	}

	saveToPNG(fo, img)
	fo.Close()	
}

// trace takes a ray and returns a colour
func trace(ray Line) NRGBA{
	
	return color.NRGBA{uint8(0), uint8(0), uint8(140), uint8(255)} 	
}

func saveToPNG(w io.Writer, img image.Image) error {
	return png.Encode(w, img)
}

