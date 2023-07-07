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
	"regexp"
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

func SplitHashtags(text string) []string {
	var (
		isHash  bool
		current string
		result  []string
	)

	for _, character := range text {
		if character == '#' {
			if current != "" {
				result = append(result, current)
			}

			current = "#"
			isHash = true
		} else {
			if isHash {
				rgx := regexp.MustCompile(`[^\w]`)

				if rgx.MatchString(string(character)) {
					result = append(result, current)

					current = ""
					isHash = false
				}
			}

			current += string(character)
		}
	}

	if current != "" {
		result = append(result, current)
	}

	return result
}
