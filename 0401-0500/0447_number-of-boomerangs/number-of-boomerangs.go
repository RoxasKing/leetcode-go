package main

// Tags:
// Math
// Hash

func numberOfBoomerangs(P [][]int) int {
	out := 0
	for _, a := range P {
		freq := make(map[int]int, len(P)-1)
		for _, b := range P {
			freq[(a[0]-b[0])*(a[0]-b[0])+(a[1]-b[1])*(a[1]-b[1])]++
		}
		for _, cnt := range freq {
			out += cnt * (cnt - 1)
		}
	}
	return out
}
