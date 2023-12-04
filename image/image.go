package image

import (
	"image"
	"image/color/palette"
	"image/gif"
	"image/png"
	"os"

	"golang.org/x/image/draw"
)

func NewRGBA(l, r int, fill func(*image.RGBA)) *image.RGBA {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{l, r}})
	fill(img)
	return img
}

func NewPaletted(l, r int, fill func(*image.Paletted), scaleFactor int) *image.Paletted {
	img := image.NewPaletted(image.Rectangle{image.Point{0, 0}, image.Point{l, r}}, palette.Plan9)
	fill(img)
	scaledRect := image.Rectangle{image.Point{0, 0}, image.Point{scaleFactor * img.Bounds().Dx(), scaleFactor * img.Bounds().Dy()}}
	scaledImg := image.NewPaletted(scaledRect, palette.Plan9)
	draw.NearestNeighbor.Scale(scaledImg, scaledRect, img, img.Bounds(), draw.Over, nil)
	return scaledImg
}

func Save(img image.Image, name string, scaleFactor int) error {
	if scaleFactor > 1 {
		scaledRect := image.Rectangle{image.Point{0, 0}, image.Point{scaleFactor * img.Bounds().Dx(), scaleFactor * img.Bounds().Dy()}}
		scaledImg := image.NewRGBA(scaledRect)
		draw.NearestNeighbor.Scale(scaledImg, scaledRect, img, img.Bounds(), draw.Over, nil)
		img = scaledImg
	}
	f, _ := os.Create(name)
	defer f.Close()
	return png.Encode(f, img)
}

func SaveGIF(frames []*image.Paletted, name string, scaleFactor, delay int) error {
	f, _ := os.Create(name)
	defer f.Close()
	delays := make([]int, len(frames))
	for i := range delays {
		delays[i] = delay
	}
	g := &gif.GIF{
		Image: frames,
		Delay: delays,
	}
	return gif.EncodeAll(f, g)
}
