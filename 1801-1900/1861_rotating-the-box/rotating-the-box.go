package main

/*
  You are given an m x n matrix of characters box representing a side-view of a box. Each cell of the box is one of the following:
    1. A stone '#'
    2. A stationary obstacle '*'
    3. Empty '.'

  The box is rotated 90 degrees clockwise, causing some of the stones to fall due to gravity. Each stone falls down until it lands on an obstacle, another stone, or the bottom of the box. Gravity does not affect the obstacles' positions, and the inertia from the box's rotation does not affect the stones' horizontal positions.

  It is guaranteed that each stone in box rests on an obstacle, another stone, or the bottom of the box.

  Return an n x m matrix representing the box after the rotation described above.

  Example 1:
    Input: box = [["#",".","#"]]
    Output: [["."],
             ["#"],
             ["#"]]

  Example 2:
    Input: box = [["#",".","*","."],
                  ["#","#","*","."]]
    Output: [["#","."],
             ["#","#"],
             ["*","*"],
             [".","."]]

  Example 3:
    Input: box = [["#","#","*",".","*","."],
                  ["#","#","#","*",".","."],
                  ["#","#","#",".","#","."]]
    Output: [[".","#","#"],
             [".","#","#"],
             ["#","#","*"],
             ["#","*","."],
             ["#",".","*"],
             ["#",".","."]]

  Constraints:
    1. m == box.length
    2. n == box[i].length
    3. 1 <= m, n <= 500
    4. box[i][j] is either '#', '*', or '.'.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/rotating-the-box
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func rotateTheBox(box [][]byte) [][]byte {
	m, n := len(box), len(box[0])
	out := make([][]byte, n)
	for i := range out {
		out[i] = make([]byte, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			out[j][m-1-i] = box[i][j]
		}
	}
	for j := 0; j < m; j++ {
		idx := n - 1
		for i := n - 1; i >= 0 && idx >= 0; i-- {
			if out[i][j] == '#' {
				out[idx][j] = '#'
				if idx != i {
					out[i][j] = '.'
				}
				idx--
			} else if out[i][j] == '*' {
				idx = i - 1
			}
		}
	}
	return out
}
