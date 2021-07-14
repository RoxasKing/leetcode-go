package main

func getConcatenation(nums []int) []int {
	n := len(nums)
	out := make([]int, n<<1)
	copy(out[:n], nums)
	copy(out[n:], nums)
	return out
}
