package main

func getMinDistance(nums []int, target int, start int) int {
	out := -1
	for i, num := range nums {
		if num == target && (out == -1 || Abs(i-start) < Abs(out-start)) {
			out = i
		}
	}
	return Abs(out - start)
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
