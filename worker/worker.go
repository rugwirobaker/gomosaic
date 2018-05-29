package worker

import (
	"image"
)

//Resize takes an image object and scales it to the given dimensions(w, h), using
//the Nearest Neighbor scaling algorithm
func Resize(src image.Image, newDims []float32) (image.Image, error) {
	b := src.Bounds()
	srcWidth := b.Dx()
	srcHeight := b.Dy()
	Xfactor := newDims[0] / float32(srcWidth)
	Yfactor := newDims[1] / float32(srcHeight)
	//create target image
	target := image.NewRGBA(image.Rect(int(
		float32(b.Min.X)*Xfactor),
		int(float32(b.Min.Y)*Yfactor),
		int(float32(b.Max.X)*Xfactor),
		int(float32(b.Max.Y)*Yfactor)))
	return target, nil
}

//int32
//func Rect(x0, y0, x1, y1 int) Rectangle {}
