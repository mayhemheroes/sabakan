package fuzz

import "github.com/cybozu-go/sabakan/v2"

func mayhemit(bytes []byte) int {

	content := int(bytes)
	//var test sabakan.LeaseRange
	//_ = sabakan.IsValidImageID(content)
	//test.IP(content)
	_ = sabakan.IsValidKernelParams(content)
	return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}