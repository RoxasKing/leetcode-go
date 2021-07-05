package main

func minCount(coins []int) int {
	out := 0
	for _, coin := range coins {
		out += coin>>1 + coin&1
	}
	return out
}
