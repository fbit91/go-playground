package main

import (
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	destinationPath := "/Users/fabriziabinetti/Downloads/AnimatedMedicalHorseshoebat_resized.gif"
	originPath := "/Users/fabriziabinetti/Downloads/AnimatedMedicalHorseshoebat.gif"

	//destinationPath := "/Users/fabriziabinetti/Downloads/Manifold_Garden_-_World_015_Stepwell_Angle_Looping_Hi-Res_resized.gif"
	//originPath := "/Users/fabriziabinetti/Downloads/Manifold_Garden_-_World_015_Stepwell_Angle_Looping_Hi-Res.gif"

	out, err := os.Create(destinationPath)
	if err != nil {
		log.Println(err)
	}
	defer out.Close()

	gifImg, err := Resize(originPath, 1980, 0)
	// write new image to file
	f, err := os.Create(destinationPath)
	defer f.Close()
	if err != nil {
		log.Println(err)
	}

	gif.EncodeAll(f, gifImg)
}

func Resize(srcFile string, width int, height int) (*gif.GIF, error) {

	f, err := os.Open(srcFile)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	im, err := gif.DecodeAll(f)

	if err != nil {
		return nil, err
	}

	if width == 0 {
		width = int(im.Config.Width * height / im.Config.Width)
	} else if height == 0 {
		height = int(width * im.Config.Height / im.Config.Width)
	}

	// reset the gif width and height
	im.Config.Width = width
	im.Config.Height = height

	firstFrame := im.Image[0].Bounds()
	img := image.NewRGBA(image.Rect(0, 0, firstFrame.Dx(), firstFrame.Dy()))

	// resize frame by frame
	for index, frame := range im.Image {
		b := frame.Bounds()
		draw.Draw(img, b, frame, b.Min, draw.Over)
		im.Image[index] = ImageToPaletted(resize.Resize(uint(width), uint(height), img, resize.NearestNeighbor))
	}

	return im, nil
}
func ImageToPaletted(img image.Image) *image.Paletted {
	b := img.Bounds()
	pm := image.NewPaletted(b, palette.WebSafe)
	draw.FloydSteinberg.Draw(pm, b, img, image.ZP)
	return pm
}
