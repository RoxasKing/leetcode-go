package main

// Tags:
// Hash
func fourSumCount(A []int, B []int, C []int, D []int) int {
	sumAB := make(map[int]int)
	for _, a := range A {
		for _, b := range B {
			sumAB[a+b]++
		}
	}
	out := 0
	for _, c := range C {
		for _, d := range D {
			out += sumAB[0-(c+d)]
		}
	}
	return out
}
