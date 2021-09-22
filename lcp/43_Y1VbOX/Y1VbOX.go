package main

// Tags:
// Dynamic Programming
// Greedy

var dirID = map[byte]int{'E': 0, 'S': 1, 'W': 2, 'N': 3}
var dirs = [][]int{
	{0, 1, 2, 3},
	{0, 1, 2}, {0, 1, 3}, {0, 2, 3}, {1, 2, 3},
	{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3},
	{0}, {1}, {2}, {3},
}
var conflicts = map[int][]int{
	01: {12, 13, 20, 21, 30, 31},
	02: {12, 13, 23, 30, 31, 32},
	03: {13, 23},
	10: {20, 30},
	12: {01, 02, 20, 23, 31, 32},
	13: {01, 02, 03, 20, 23, 30},
	20: {01, 10, 12, 13, 30, 31},
	21: {01, 31},
	23: {02, 03, 12, 13, 30, 31},
	30: {01, 02, 10, 13, 20, 23},
	31: {01, 02, 12, 20, 21, 23},
	32: {02, 12},
}

type State struct {
	t, c int
	p    [4]int
}

func trafficCommand(directions []string) int {
	directionArr := [4][]int{}
	for i := range directions {
		for j := range directions[i] {
			directionArr[i] = append(directionArr[i], dirID[directions[i][j]])
		}
	}
	isConflict := [33][33]bool{}
	for x, arr := range conflicts {
		for _, y := range arr {
			isConflict[x][y] = true
		}
	}
	cnt := 0
	for _, dir := range directions {
		cnt += len(dir)
	}
	minT := [21][21][21][21]int{}
	for i := 0; i < 21; i++ {
		for j := 0; j < 21; j++ {
			for k := 0; k < 21; k++ {
				for l := 0; l < 21; l++ {
					minT[i][j][k][l] = 1<<31 - 1
				}
			}
		}
	}
	states := []*State{{t: 0, c: cnt, p: [4]int{}}}
	for ; len(states) > 0; states = states[1:] {
		state := states[0]
		t, c, p := state.t, state.c, state.p
		if minT[p[0]][p[1]][p[2]][p[3]] <= t {
			continue
		}
		minT[p[0]][p[1]][p[2]][p[3]] = t
		mark := [15]bool{}
		for k, dir := range dirs {
			if mark[0] ||
				mark[1] && (k == 5 || k == 6 || k == 8 || k == 11 || k == 12 || k == 13) ||
				mark[2] && (k == 5 || k == 7 || k == 9 || k == 11 || k == 12 || k == 14) ||
				mark[3] && (k == 6 || k == 7 || k == 10 || k == 11 || k == 13 || k == 14) ||
				mark[4] && (k == 8 || k == 9 || k == 10 || k == 12 || k == 13 || k == 14) ||
				mark[5] && (k == 11 || k == 12) ||
				mark[6] && (k == 11 || k == 13) ||
				mark[7] && (k == 11 || k == 14) ||
				mark[8] && (k == 12 || k == 13) ||
				mark[9] && (k == 12 || k == 14) ||
				mark[10] && (k == 13 || k == 14) {
				continue
			}
			ok := true
			arr := make([]int, 0, 4)
			for _, i := range dir {
				if p[i] == len(directionArr[i]) {
					ok = false
					break
				}
				arr = append(arr, i*10+directionArr[i][p[i]])
			}
			for i := 0; i < len(arr) && ok; i++ {
				for j := i + 1; j < len(arr) && ok; j++ {
					if isConflict[arr[i]][arr[j]] {
						ok = false
					}
				}
			}
			if !ok {
				continue
			}
			if c-len(dir) == 0 {
				return t + 1
			}
			mark[k] = true
			newState := &State{c: c - len(dir), t: t + 1, p: state.p}
			for _, i := range dir {
				newState.p[i]++
			}
			states = append(states, newState)
		}
	}
	return cnt
}
