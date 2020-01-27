package rcrn

import (
	"fmt"
)

func Profile() {
	fmt.Println("Generating a million bits")
	var bits = RandBits(1_000_000)
	fmt.Println("Getting distribution")
	var zero = 0
	for i := 0; i < len(bits); i++ {
		if bits[i] == 0 {
			zero++
		}
	}
	fmt.Printf("%v%% zero", (float64(zero)/float64(len(bits))) * 100)
}
