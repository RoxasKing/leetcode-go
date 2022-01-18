package main

// Difficulty:
// Medium

func maxDistToClosest(seats []int) int {
	var out, cnt int
	for i, seat := range seats {
		if seat == 0 {
			cnt++
		} else {
			if cnt == i {
				out = Max(out, cnt)
			} else {
				out = Max(out, (cnt+1)>>1)
			}
			cnt = 0
		}
	}
	if seats[len(seats)-1] == 0 {
		out = Max(out, cnt)
	} else {
		out = Max(out, (cnt+1)>>1)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
