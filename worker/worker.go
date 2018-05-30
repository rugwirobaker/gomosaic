package worker

import (
	"image"
	"image/color"
)

//Resize takes an image object and scales it to the given dimensions(w, h), using
//the Nearest Neighbor scaling algorithm
func Resize(src image.Image, newDims []float32) image.Image {
	b := src.Bounds()
	srcWidth := b.Dx()
	srcHeight := b.Dy()
	Xfactor := newDims[0] / float32(srcWidth)
	Yfactor := newDims[1] / float32(srcHeight)
	//create target image
	target := image.NewNRGBA(image.Rect(int(
		float32(b.Min.X)*Xfactor),
		int(float32(b.Min.Y)*Yfactor),
		int(float32(b.Max.X)*Xfactor),
		int(float32(b.Max.Y)*Yfactor)),
	)
	for y, j := b.Min.Y, b.Min.Y; j < b.Max.Y; y, j = y+int(Yfactor),
		j+1 {
		for x, i := b.Min.X, b.Min.X; i < b.Max.X; x, i = x+int(Xfactor),
			i+1 {
			r, g, b, a := src.At(x, y).RGBA()
			target.SetNRGBA(i, j, color.NRGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8),
				uint8(a >> 8)})
		}
	}
	return target
}

//AverageColor takes an image object as
func AverageColor(img image.Image) [3]float64 {
	bounds := img.Bounds()
	r, g, b := 0.0, 0.0, 0.0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1)
		}
	}
	tPixels := float64(bounds.Max.X * bounds.Max.Y)
	return [3]float64{r / tPixels, g / tPixels, b / tPixels}
}
