package main

// Tags:
// Hash

func tupleSameProduct(nums []int) int {
	n := len(nums)
	cnt := map[int]int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			cnt[nums[i]*nums[j]]++
		}
	}
	out := 0
	for _, c := range cnt {
		out += 4 * c * (c - 1)
	}
	return out
}
