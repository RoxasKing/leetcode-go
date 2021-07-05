package main

// Tags:
// Prefix Sum + Hash
func subarraySum(nums []int, k int) int {
	out := 0
	pSum := make(map[int]int)
	pSum[0] = 1
	sum := 0
	for _, num := range nums {
		sum += num
		if cnt, ok := pSum[sum-k]; ok {
			out += cnt
		}
		pSum[sum]++
	}
	return out
}
