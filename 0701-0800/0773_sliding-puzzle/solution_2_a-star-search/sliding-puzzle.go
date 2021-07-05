package main

import (
	"container/heap"
)

// Tags:
// A* Search Algorithm
// Priority Queue

func slidingPuzzle(board [][]int) int {
	state := [2][3]int{{board[0][0], board[0][1], board[0][2]}, {board[1][0], board[1][1], board[1][2]}}
	visited := [6 * 6 * 6 * 6 * 6 * 6]bool{}
	visited[getIdx(state)] = true
	h := MinHeap{}
	heap.Push(&h, &elem{state: state, g: 0, h: getDist(state)})
	out := -1

	for h.Len() > 0 {
		e := heap.Pop(&h).(*elem)
		state, moves := e.state, e.g
		if getDist(state) == 0 {
			out = moves
			break
		}

		x0, y0 := 0, 0
		for i := 0; i < 6; i++ {
			x, y := i/3, i%3
			if state[x][y] == 0 {
				x0, y0 = x, y
				break
			}
		}

		for _, fwd := range fwds {
			x, y := x0+fwd[0], y0+fwd[1]
			if x < 0 || 1 < x || y < 0 || 2 < y {
				continue
			}

			state[x0][y0], state[x][y] = state[x][y], state[x0][y0]
			idx := getIdx(state)
			if !visited[idx] {
				visited[idx] = true
				heap.Push(&h, &elem{state: state, g: moves + 1, h: getDist(state)})
			}
			state[x0][y0], state[x][y] = state[x][y], state[x0][y0]
		}
	}
	return out
}

var fwds = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

type elem struct {
	g, h  int
	state [2][3]int
}

type MinHeap []*elem

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].g+h[i].h < h[j].g+h[j].h }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*elem)) }
func (h *MinHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}

func getDist(board [2][3]int) int {
	out := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			x, y := getCoordinate(board[i][j])
			out += Abs(x-i) + Abs(y-j)
		}
	}
	return out
}

var dict = [][2]int{
	1: {0, 0},
	2: {0, 1},
	3: {0, 2},
	4: {1, 0},
	5: {1, 1},
	0: {1, 2},
}

func getCoordinate(num int) (int, int) {
	res := dict[num]
	return res[0], res[1]
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func getIdx(state [2][3]int) int {
	a, b, c, d, e, f := state[0][0], state[0][1], state[0][2], state[1][0], state[1][1], state[1][2]
	return ((((a*6+b)*6+c)*6+d)*6+e)*6 + f
}
