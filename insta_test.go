package insta_go

import "testing"

func TestBuildPost(t *testing.T) {
	avatar := "https://cdn.discordapp.com/avatars/669523636423622686/1a3b7e67364d356960bd07e85c8a00f1.png?size=4096"
	image := "https://media.tenor.com/2vtPJ3IkKmMAAAAd/selfie-cat.gif"

	post, err := BuildPost(avatar, "Laura", image, "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut")
	if err != nil {
		t.Fatal(err)
	}

	err = SavePNG(post, "test.png")
	if err != nil {
		t.Fatal(err)
	}
}
