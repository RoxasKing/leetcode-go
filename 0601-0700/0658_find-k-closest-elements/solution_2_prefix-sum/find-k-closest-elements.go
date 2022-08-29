package main

// Difficulty:
// Medium

// Tags:
// Sliding Window

func findClosestElements(arr []int, k int, x int) []int {
	o, min, cur := -1, 1<<31-1, 0
	for i, num := range arr {
		cur += abs(num - x)
		if i > k-1 {
			cur -= abs(arr[i-k] - x)
		}
		if i >= k-1 && min > cur {
			o, min = i-k+1, cur
		}
	}
	return arr[o : o+k]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
