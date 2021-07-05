package main

// Tags:
// Simulation

func findLatestStep(arr []int, m int) int {
	n := len(arr)
	gCnt := make([]int, n+1)
	gLen := make([]int, n)
	out := -1

	update := func(l, r int) {
		gCnt[gLen[l]]--
		gCnt[gLen[r]]--
		newLen := gLen[l] + gLen[r]
		gCnt[newLen]++
		gLen[l] = newLen
		gLen[r] = newLen
	}

	for j, i := range arr {
		i--
		j++
		gCnt[1]++
		gLen[i] = 1
		if i-1 >= 0 && gLen[i-1] > 0 {
			l := i - 1 - gLen[i-1] + 1
			r := i - gLen[i] + 1
			update(l, r)
		}
		if i+1 < n && gLen[i+1] > 0 {
			l := i - gLen[i] + 1
			r := i + 1 + gLen[i+1] - 1
			update(l, r)
		}
		if gCnt[m] > 0 {
			out = j
		}
	}

	return out
}
