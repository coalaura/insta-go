package insta_go

import (
	"github.com/nfnt/resize"
	"image"
	"image/draw"
)

func BuildPost(avatar, username, post, text string) (image.Image, error) {
	header, err := BuildHeader(username, avatar)
	if err != nil {
		return nil, err
	}

	footer, err := BuildFooter(text)
	if err != nil {
		return nil, err
	}

	postImage, err := DownloadImage(post)
	if err != nil {
		return nil, err
	}

	scaledImage := _fitPost(postImage, header.Bounds().Max.X-6)

	headerTop := header.Bounds().Max.Y
	postTop := scaledImage.Bounds().Max.Y
	footerTop := footer.Bounds().Max.Y

	combined := image.NewRGBA(image.Rect(0, 0, header.Bounds().Max.X, headerTop+postTop+footerTop))

	draw.Draw(combined, combined.Bounds(), image.White, image.Point{}, draw.Src)

	draw.Draw(combined, header.Bounds(), header, image.Point{}, draw.Over)
	draw.Draw(combined, scaledImage.Bounds().Add(image.Point{X: 3, Y: headerTop}), scaledImage, image.Point{}, draw.Over)
	draw.Draw(combined, footer.Bounds().Add(image.Point{X: 0, Y: headerTop + postTop}), footer, image.Point{}, draw.Over)

	return combined, nil
}

func _fitPost(post image.Image, width int) image.Image {
	ratio := float64(post.Bounds().Max.X) / float64(width)

	height := int(float64(post.Bounds().Max.Y) / ratio)

	return resize.Resize(uint(width), uint(height), post, resize.Lanczos3)
}
