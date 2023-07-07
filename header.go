package insta_go

import (
	_ "embed"
	"image"
	"image/draw"

	"github.com/nfnt/resize"
)

var (
	//go:embed assets/header.png
	headerBytes []byte

	avatarPosition = image.Point{X: 20, Y: 20}
	avatarSize     = image.Point{X: 78, Y: 78}

	usernamePosition = image.Point{X: 120, Y: 67}
)

func BuildHeader(username, avatar string) (image.Image, error) {
	header, err := LoadPNG(headerBytes)
	if err != nil {
		return nil, err
	}

	avatarImage, err := DownloadImage(avatar)
	if err != nil {
		return nil, err
	}

	scaledAvatar := resize.Resize(uint(avatarSize.X), uint(avatarSize.Y), avatarImage, resize.Lanczos3)

	combined := image.NewRGBA(header.Bounds())

	draw.Draw(combined, scaledAvatar.Bounds().Add(avatarPosition), scaledAvatar, image.Point{}, draw.Src)
	draw.Draw(combined, header.Bounds(), header, image.Point{}, draw.Over)

	pos := usernamePosition

	err = WriteText(combined, &pos, 26, username, image.Black)
	if err != nil {
		return nil, err
	}

	return combined, nil
}
