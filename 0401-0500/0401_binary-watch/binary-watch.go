package main

import (
	"fmt"
	"math/bits"
)

// Tags:
// Bit Manipulation

func readBinaryWatch(turnedOn int) []string {
	out := []string{}
	for h := uint8(0); h < 12; h++ {
		for m := uint8(0); m < 60; m++ {
			if bits.OnesCount8(h)+bits.OnesCount8(m) == turnedOn {
				out = append(out, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return out
}
