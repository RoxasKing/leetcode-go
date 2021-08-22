package main

func escapeGhosts(ghosts [][]int, target []int) bool {
	d := Abs(target[0]) + Abs(target[1])
	for _, g := range ghosts {
		if Abs(g[0]-target[0])+Abs(g[1]-target[1]) <= d {
			return false
		}
	}
	return true
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
