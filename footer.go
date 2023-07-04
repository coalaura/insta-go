package main

import (
	_ "embed"
	"image"
	"image/draw"
	"strings"
)

var (
	//go:embed assets/footer.png
	footerBytes []byte

	footerSize     = image.Point{X: 0, Y: 48}
	footerPosition = image.Point{X: 28, Y: 20}
)

func BuildFooter(text string) (image.Image, error) {
	footer, err := LoadPNG(footerBytes)
	if err != nil {
		return nil, err
	}

	if text == "" {
		return footer, nil
	}

	text, err = _capText(text)
	if err != nil {
		return nil, err
	}

	size := footer.Bounds()
	size.Max.Y += footerSize.Y

	combined := image.NewRGBA(size)

	draw.Draw(combined, combined.Bounds(), image.White, image.Point{}, draw.Src)
	draw.Draw(combined, footer.Bounds(), footer, image.Point{}, draw.Over)

	pos := footerPosition.Add(image.Point{X: 0, Y: footer.Bounds().Max.Y})

	err = WriteText(combined, pos, 26, text)
	if err != nil {
		return nil, err
	}

	return combined, nil
}

func _capText(text string) (string, error) {
	first := true

	text = strings.TrimSpace(text)

	text = strings.Split(text, "\n")[0]

	for {
		l, err := MeasureText(text, 26)
		if err != nil {
			return "", err
		}

		if l <= 920 {
			if !first {
				text += "..."
			}

			return text, nil
		} else {
			text = text[:len(text)-1]
		}

		first = false
	}
}
