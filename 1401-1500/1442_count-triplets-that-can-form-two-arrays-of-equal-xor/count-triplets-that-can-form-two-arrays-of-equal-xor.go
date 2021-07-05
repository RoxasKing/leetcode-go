package main

// Tags:
// Bit Manipulation

func countTriplets(arr []int) int {
	out := 0
	n := len(arr)
	for i := range arr {
		xor := arr[i]
		for j := i + 1; j < n; j++ {
			xor ^= arr[j]
			if xor == 0 {
				out += j - i
			}
		}
	}
	return out
}
