package main

// Difficulty:
// Easy

func selfDividingNumbers(left int, right int) []int {
	out := []int{}
	for num := left; num <= right; num++ {
		ok := true
		for i := num; i > 0; i /= 10 {
			if x := i % 10; x == 0 || num%x != 0 {
				ok = false
				break
			}
		}
		if ok {
			out = append(out, num)
		}
	}
	return out
}
