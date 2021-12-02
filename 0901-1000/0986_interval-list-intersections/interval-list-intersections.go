package main

// Difficulty:
// Medium

// Tags:
// Two Pointers

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	m, n := len(firstList), len(secondList)
	out := [][]int{}
	for i, j := 0, 0; i < m && j < n; {
		if firstList[i][1] < secondList[j][0] {
			i++
		} else if firstList[i][0] > secondList[j][1] {
			j++
		} else {
			l := Max(firstList[i][0], secondList[j][0])
			r := Min(firstList[i][1], secondList[j][1])
			if firstList[i][1] == r {
				i++
			} else {
				j++
			}
			out = append(out, []int{l, r})
		}
	}
	return out
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
