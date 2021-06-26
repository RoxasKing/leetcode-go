package main

/*
  On a 2x3 board, there are 5 tiles represented by the integers 1 through 5, and an empty square represented by 0.

  A move consists of choosing 0 and a 4-directionally adjacent number and swapping it.

  The state of the board is solved if and only if the board is [[1,2,3],[4,5,0]].

  Given a puzzle board, return the least number of moves required so that the state of the board is solved. If it is impossible for the state of the board to be solved, return -1.

  Examples:
    Input: board = [[1,2,3],[4,0,5]]
    Output: 1
    Explanation: Swap the 0 and the 5 in one move.

    Input: board = [[1,2,3],[5,4,0]]
    Output: -1
    Explanation: No number of moves will make the board solved.

    Input: board = [[4,1,2],[5,0,3]]
    Output: 5
    Explanation: 5 is the smallest number of moves that solves the board.
      An example path:
      After move 0: [[4,1,2],[5,0,3]]
      After move 1: [[4,1,2],[0,5,3]]
      After move 2: [[0,1,2],[4,5,3]]
      After move 3: [[1,0,2],[4,5,3]]
      After move 4: [[1,2,0],[4,5,3]]
      After move 5: [[1,2,3],[4,5,0]]
      Input: board = [[3,2,4],[1,5,0]]
      Output: 14

  Note:
    1. board will be a 2 x 3 array as described above.
    2. board[i][j] will be a permutation of [0, 1, 2, 3, 4, 5].

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sliding-puzzle
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// BFS

func slidingPuzzle(board [][]int) int {
	target := ((((1*6+2)*6+3)*6+4)*6 + 5) * 6
	visited := [6 * 6 * 6 * 6 * 6 * 6]bool{}
	visited[getIdx(board[0][0], board[0][1], board[0][2], board[1][0], board[1][1], board[1][2])] = true
	q := [][7]int{{board[0][0], board[0][1], board[0][2], board[1][0], board[1][1], board[1][2]}}
	out := -1
	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		cur, t := getIdx(e[0], e[1], e[2], e[3], e[4], e[5]), e[6]
		if cur == target {
			out = t
			break
		}
		a, x0, y0 := 0, 0, 0
		for i := 0; i < 6; i++ {
			if e[i] == 0 {
				a, x0, y0 = i, i/3, i%3
				break
			}
		}
		for _, fwd := range fwds {
			x, y := x0+fwd[0], y0+fwd[1]
			if x < 0 || 1 < x || y < 0 || 2 < y {
				continue
			}
			b := x*3 + y
			ne := e
			ne[a], ne[b], ne[6] = ne[b], ne[a], t+1
			idx := getIdx(ne[0], ne[1], ne[2], ne[3], ne[4], ne[5])
			if visited[idx] {
				continue
			}
			visited[idx] = true
			q = append(q, ne)
		}
	}
	return out
}

var fwds = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func getIdx(a, b, c, d, e, f int) int {
	return ((((a*6+b)*6+c)*6+d)*6+e)*6 + f
}
