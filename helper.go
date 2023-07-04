package insta_go

import (
	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
)

func LoadPNG(b []byte) (image.Image, error) {
	return png.Decode(bytes.NewReader(b))
}

func DownloadImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	return img, nil
}

func SavePNG(img image.Image, path string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		return err
	}

	return nil
}
