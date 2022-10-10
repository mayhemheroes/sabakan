package fuzz

import "github.com/cybozu-go/sabakan/v2"

func mayhemit(bytes []byte) int {

	content := string(bytes)
	var test sabakan.ImageIndex
	//_ = sabakan.IsValidImageID(content)
	test.Remove(content)
	return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}