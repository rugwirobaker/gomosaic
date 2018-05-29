package worker

import (
	"image"
	"testing"
)

func TestResize(t *testing.T) {
	testScale := [][]float32{
		//{width, height}
		{5, 10},
		{10, 15},
		{15, 20},
		{50, 50},
		{60, 75},
		{85, 95},
		{100, 110},
		{105, 115},
		{108, 119},
	}
	testImage := []image.Image{
		image.NewRGBA64(image.Rect(0, 0, 50, 50)),
		image.NewNRGBA64(image.Rect(0, 0, 50, 50)),
		image.NewAlpha16(image.Rect(0, 0, 50, 50)),
		image.NewGray16(image.Rect(0, 0, 50, 50)),
	}
	img, err := Resize(testImage[0], testScale[0])
	//Assertion1: test if Resize returns a image
	if img == nil && err != nil {
		t.Error("Expected an image output got", err)
		return
	}
	//Assertion1: test if given image has the given new scales
	for _, i := range testImage {
		for _, s := range testScale {
			img, err = Resize(i, s)
			//if img.Scales != s then print an error[expected the img to be of size s]
			imgSize := []int{
				img.Bounds().Dx(), //width
				img.Bounds().Dy(), //height
			}
			if imgSize[0] != int(s[0]) && imgSize[1] != int(s[1]) {
				t.Error("Expected image to be of size", s[0], s[1],
					"but got", imgSize[0], imgSize[1])
			}
		}
	}
}

//func TestAverageColor(t *testing.T) {}

//func TestLoadTilesRepo(t *testing.T) {}

//func TestDecodeImage(t *testing.T) {}
