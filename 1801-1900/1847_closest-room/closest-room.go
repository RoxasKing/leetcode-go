package main

import (
	"sort"
)

func closestRoom(rooms [][]int, queries [][]int) []int {
	n := len(queries)

	idxs := make([]int, n)
	for i := range idxs {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool { return queries[idxs[i]][1] < queries[idxs[j]][1] })
	sort.Slice(rooms, func(i, j int) bool { return rooms[i][1] < rooms[j][1] })

	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = -1
	}

	for _, idx := range idxs {
		q := queries[idx]
		i := sort.Search(len(rooms), func(i int) bool { return rooms[i][1] >= q[1] })
		rooms = rooms[i:]
		for _, room := range rooms {
			if out[idx] == -1 || Abs(out[idx]-q[0]) > Abs(room[0]-q[0]) ||
				Abs(out[idx]-q[0]) == Abs(room[0]-q[0]) && room[0] < out[idx] {
				out[idx] = room[0]
			}
			if Abs(out[idx]-q[0]) == 0 {
				break
			}
		}
	}

	return out
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
