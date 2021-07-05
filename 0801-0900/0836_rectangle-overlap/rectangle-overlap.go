package main

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return Max(rec1[0], rec2[0]) < Min(rec1[2], rec2[2]) && Max(rec1[1], rec2[1]) < Min(rec1[3], rec2[3])
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
