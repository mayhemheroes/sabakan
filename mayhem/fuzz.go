package fuzz

import "github.com/cybozu-go/sabakan/v2"

func mayhemit(bytes []byte) int {

	content := string(bytes)
	_ = sabakan.IsValidImageID(content)
	return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}