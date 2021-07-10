package main

// Tags:
// Hash

func numSubarraysWithSum(nums []int, goal int) int {
	var out, sum int
	dict := map[int]int{0: 1}
	for _, num := range nums {
		sum += num
		if cnt, ok := dict[sum-goal]; ok {
			out += cnt
		}
		dict[sum]++
	}
	return out
}
