package main

// Difficulty:
// Hard

// Tags:
// DFS
// Memorization

const (
	draw = iota
	mouseWin
	catWin
)

func catMouseGame(graph [][]int) int {
	n := len(graph)
	f := make([][][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			f[i][j] = make([]int, n<<1)
			for k := 0; k < n<<1; k++ {
				f[i][j][k] = -1
			}
		}
	}
	var getResult, getNextResult func(int, int, int) int
	getResult = func(mouse, cat, turns int) int {
		if turns == n<<1 {
			return draw
		}
		out := f[mouse][cat][turns]
		if out != -1 {
			return out
		}
		if mouse == 0 {
			out = mouseWin
		} else if mouse == cat {
			out = catWin
		} else {
			out = getNextResult(mouse, cat, turns)
		}
		f[mouse][cat][turns] = out
		return out
	}
	getNextResult = func(mouse, cat, turns int) int {
		curMove := mouse
		if turns&1 == 1 {
			curMove = cat
		}
		defaultOut := mouseWin
		if curMove == mouse {
			defaultOut = catWin
		}
		out := defaultOut
		for _, next := range graph[curMove] {
			if curMove == cat && next == 0 {
				continue
			}
			nextMouse, nextCat := mouse, cat
			if curMove == mouse {
				nextMouse = next
			} else if curMove == cat {
				nextCat = next
			}
			nextOut := getResult(nextMouse, nextCat, turns+1)
			if nextOut != defaultOut {
				if out = nextOut; out != draw {
					break
				}
			}
		}
		return out
	}
	return getResult(1, 2, 0)
}
