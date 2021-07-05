package main

func maximumPopulation(logs [][]int) int {
	min, max := 1<<31-1, 0
	for _, log := range logs {
		min = Min(min, log[0])
		max = Max(max, log[1])
	}

	out, cnt := 0, 0
	for i := min; i <= max; i++ {
		tmp := 0
		for _, log := range logs {
			if log[0] <= i && i <= log[1]-1 {
				tmp++
			}
		}
		if tmp > cnt {
			out, cnt = i, tmp
		}
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
