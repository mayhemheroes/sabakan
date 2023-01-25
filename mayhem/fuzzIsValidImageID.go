package fuzzIsValidImageID

import "strconv"
import "github.com/cybozu-go/sabakan/v2"
import fuzz "github.com/AdaLogics/go-fuzz-headers"

func mayhemit(bytes []byte) int {

	var num int
    if len(bytes) > 2 {
        num, _ = strconv.Atoi(string(bytes[0]))
        bytes = bytes[1:]

        switch num {
    
        case 0:
			fuzzConsumer := fuzz.NewConsumer(bytes)
			var content string
			err := fuzzConsumer.CreateSlice(&content)
			if err != nil {
					return 0
			}

			_ = sabakan.IsValidImageID(content)
			return 0

		default:
			fuzzConsumer := fuzz.NewConsumer(bytes)
			var content string
			err := fuzzConsumer.CreateSlice(&content)
			if err != nil {
					return 0
			}

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