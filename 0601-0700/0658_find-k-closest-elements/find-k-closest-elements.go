package main

// Tags:
// Sliding Window
func findClosestElements(arr []int, k int, x int) []int {
	out := make([]int, 0, k)
	for i := range arr {
		if len(out) == k && Abs(arr[i]-x) < Abs(out[0]-x) {
			out = out[1:]
			out = append(out, arr[i])
		} else if len(out) < k {
			out = append(out, arr[i])
		}
	}
	return out
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
