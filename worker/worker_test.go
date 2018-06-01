package worker

import (
	"encoding/base64"
	"image"
	_ "image/jpeg"
	"reflect"
	"strings"
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
	img := Resize(testImage[0], testScale[0])
	//Assertion1: test if Resize returns a image
	if img == nil {
		t.Error("Expected an image output got", img)
		return
	}
	//Assertion1: test if given image has the given new scales
	for _, i := range testImage {
		for _, s := range testScale {
			img = Resize(i, s)
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

	//test resize with non-empty image
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, _, err := image.Decode(reader)
	if err != nil {
		t.Error("Got error:'", err, "'while trying to decode test img")
		return
	}
	img = Resize(m, testScale[0])
	if img == nil {
		t.Error("Got error:", img, "while trying to decode test img")
		return
	}
	//tAvgColor := avgColor(m)
	iAvgColor := avgColor(img)

	if iAvgColor == [3]float64{0, 0, 0} {
		t.Error("The image does not contain any pixels:", iAvgColor)
	}
	//var diffRatio float64 // 0.50
}

func TestAverageColor(t *testing.T) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	img, _, err := image.Decode(reader)
	if err != nil {
		t.Error("Got error:'", err, "'while trying to decode test img")
		return
	}
	//test the returned type
	testAvg := AverageColor(img)

	//generate a test value
	trueAvg := avgColor(img)
	if reflect.TypeOf(testAvg).Name() != reflect.TypeOf(trueAvg).Name() {
		t.Error("Expected type: float64 got", reflect.TypeOf(testAvg).Name())
		return
	}

	//test if the value is right
	if testAvg != trueAvg {
		t.Error("Expected a value of:", trueAvg, "got", testAvg)
	}
}

//testing the tile DB
