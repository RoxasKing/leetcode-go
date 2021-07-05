package main

// Tags:
// Hash
func leastBricks(wall [][]int) int {
	cnt := make(map[int]int)
	for i := range wall {
		sum := 0
		for j := 0; j < len(wall[i])-1; j++ {
			sum += wall[i][j]
			cnt[sum]++
		}
	}

	n := len(wall)
	out := n
	for _, c := range cnt {
		out = Min(out, n-c)
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
