package main

// Tags:
// Math
func countNicePairs(nums []int) int {
	cnt := make(map[int]int)
	for i := range nums {
		cnt[nums[i]-rev(nums[i])]++
	}
	out := 0
	for _, x := range cnt {
		res := (x * (x - 1) / 2) % (1e9 + 7)
		out = (out + res) % (1e9 + 7)
	}
	return out
}

func rev(num int) int {
	out := 0
	for num > 0 {
		out = out*10 + num%10
		num /= 10
	}
	return out
}
