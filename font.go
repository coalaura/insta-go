package main

import (
	_ "embed"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
)

var (
	//go:embed assets/normal.ttf
	fontBytes []byte

	fontFace *truetype.Font
)

func _loadFont() error {
	if fontFace == nil {
		f, err := truetype.Parse(fontBytes)
		if err != nil {
			return err
		}

		fontFace = f
	}

	return nil
}

func WriteText(img *image.RGBA, pos image.Point, size float64, text string) error {
	err := _loadFont()
	if err != nil {
		return err
	}

	d := &font.Drawer{
		Dst: img,
		Src: image.Black,
		Face: truetype.NewFace(fontFace, &truetype.Options{
			Size:    size,
			DPI:     72,
			Hinting: font.HintingNone,
		}),
	}

	d.Dot = fixed.Point26_6{
		X: fixed.I(pos.X),
		Y: fixed.I(pos.Y),
	}

	d.DrawString(text)

	return nil
}

func MeasureText(text string, size float64) (int, error) {
	err := _loadFont()
	if err != nil {
		return 0, err
	}

	d := &font.Drawer{
		Face: truetype.NewFace(fontFace, &truetype.Options{
			Size:    size,
			DPI:     72,
			Hinting: font.HintingNone,
		}),
	}

	return d.MeasureString(text).Ceil(), nil
}
