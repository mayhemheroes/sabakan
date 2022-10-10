package fuzz

import "github.com/cybozu-go/sabakan/v2"

func mayhemit(bytes []byte) int {

	content := string(bytes)
	var test sabakan.ImageIndex
	var testImg sabakan.Image
	test.Append(&testImg)
	//_ = sabakan.IsValidImageID(content)
	test.Find(content)
	return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}