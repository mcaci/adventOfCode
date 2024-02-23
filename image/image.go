package image

import (
	"image"
	"image/color/palette"
	"image/gif"
	"image/png"
	"io"
	"os"

	"golang.org/x/image/draw"
)

func createRGBA(r image.Rectangle) *image.RGBA         { return image.NewRGBA(r) }
func createGifFrame(r image.Rectangle) *image.Paletted { return image.NewPaletted(r, palette.Plan9) }

func New[T any](create func(image.Rectangle) T) func(l, r int, fill func(T) T) T {
	return func(l, r int, fill func(T) T) T {
		rect := image.Rectangle{image.Point{0, 0}, image.Point{l, r}}
		img := create(rect)
		return fill(img)
	}
}

var NewImage = New(createRGBA)
var NewGifFrame = New(createGifFrame)

func Scale[T draw.Image](create func(image.Rectangle) T) func(img T, f int) T {
	return func(img T, f int) T {
		if f < 1 {
			return img
		}
		rect := image.Rectangle{image.Point{0, 0}, image.Point{f * img.Bounds().Dx(), f * img.Bounds().Dy()}}
		scaledImg := create(rect)
		draw.NearestNeighbor.Scale(scaledImg, rect, img, img.Bounds(), draw.Over, nil)
		return scaledImg
	}
}

var ScaleImage = Scale(createRGBA)
var ScaleGifFrame = Scale(createGifFrame)

func Save[T any](name string, obj T, saveF func(io.Writer, T) error) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return saveF(f, obj)
}

func SaveImg(name string, img image.Image) error {
	return Save(name, img, func(w io.Writer, img image.Image) error { return png.Encode(w, img) })
}

func SaveGIF(name string, g *gif.GIF) error {
	return Save(name, g, func(w io.Writer, g *gif.GIF) error { return gif.EncodeAll(w, g) })
}

func NewGif(frames []*image.Paletted, delay int) *gif.GIF {
	delays := make([]int, len(frames))
	for i := range delays {
		delays[i] = delay
	}
	return &gif.GIF{
		Image: frames,
		Delay: delays,
	}
}
