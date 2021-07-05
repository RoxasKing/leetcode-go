package main

// Tags:
// Math
func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	dif := goal - sum
	out := Abs(dif) / Abs(limit)
	if Abs(dif)%Abs(limit) > 0 {
		out++
	}
	return out
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
