package main

import (
	"math"
	"math/bits"
	"strconv"
)

// Tags:
// Math

func smallestGoodBase(n string) string {
	nVal, _ := strconv.Atoi(n)
	for m := bits.Len(uint(nVal)) - 1; m > 1; m-- {
		k := int(math.Pow(float64(nVal), 1/float64(m)))
		mul, sum := 1, 1
		for i := 0; i < m; i++ {
			mul *= k
			sum += mul
		}
		if sum == nVal {
			return strconv.Itoa(k)
		}
	}
	return strconv.Itoa(nVal - 1)
}
