package main

// Tags:
// Prefix Sum
// Hash

func checkSubarraySum(nums []int, k int) bool {
	dict := map[int]int{0: -1}
	sum := 0
	for j, num := range nums {
		sum += num
		sum %= k
		if i, ok := dict[sum]; !ok {
			dict[sum] = j
		} else if j-i > 1 {
			return true
		}
	}
	return false
}
