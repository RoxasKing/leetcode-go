package main

// Tags:
// TODO

// Dynamic Programming
func minimumOperations(leaves string) int {
	n := len(leaves)
	g := -1
	if leaves[0] == 'y' {
		g = 1
	}
	gmin := g
	min := 1<<31 - 1
	for i := 1; i < n-1; i++ {
		g += 2*isYellow(leaves[i] == 'y') - 1
		min = Min(min, gmin-g)
		gmin = Min(gmin, g)
	}
	if leaves[n-1] == 'y' {
		g++
	}
	return min + (g+n)/2
}

func isYellow(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
