package insta_go

import "testing"

func TestBuildPost(t *testing.T) {
	avatar := "https://i.shrt.day/vocEbIsI30.png"
	image := "https://i.shrt.day/monUFoPU69.png"

	post, err := BuildPost(avatar, "Laura", image, "Lorem ipsum dolor sit #test #cool.yeah", false)
	if err != nil {
		t.Fatal(err)
	}

	err = SavePNG(post, "test.png")
	if err != nil {
		t.Fatal(err)
	}
}
