package fuzzIsValidImageID

import "strconv"
import "github.com/cybozu-go/sabakan/v2"

func mayhemit(bytes []byte) int {

	var num int
    if len(bytes) > 2 {
        num, _ = strconv.Atoi(string(bytes[0]))
        bytes = bytes[1:]

        switch num {
    
        case 0:
			content := string(bytes)
			_ = sabakan.IsValidImageID(content)
			return 0

		default:
			content := string(bytes)
			_ = sabakan.IsValidKernelParams(content)
			return 0
		}
	}
	return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}