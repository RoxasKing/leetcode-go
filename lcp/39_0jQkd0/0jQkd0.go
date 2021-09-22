package main

func minimumSwitchingTimes(source [][]int, target [][]int) int {
	freq := [1e4 + 1]int{}
	for i := range source {
		for _, x := range source[i] {
			freq[x]++
		}
	}
	for i := range target {
		for _, x := range target[i] {
			freq[x]--
		}
	}
	var out int
	for i := 1; i <= 1e4; i++ {
		out += Abs(freq[i])
	}
	return out >> 1
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
