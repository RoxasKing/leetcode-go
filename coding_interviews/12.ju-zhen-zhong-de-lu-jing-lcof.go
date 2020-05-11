package codinginterviews

/*
  请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。
    [["a","b","c","e"],
    ["s","f","c","s"],
    ["a","d","e","e"]]
  但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。

  提示：
    1 <= board.length <= 200
    1 <= board[i].length <= 200

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func exist(board [][]byte, word string) bool {
	steps := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	var index int
	var check func(int, int) bool
	check = func(x, y int) bool {
		if index == len(word) {
			return true
		}
		for _, step := range steps {
			nx, ny := x+step[0], y+step[1]
			if nx < 0 || nx > len(board)-1 ||
				ny < 0 || ny > len(board[0])-1 ||
				visited[nx][ny] ||
				board[nx][ny] != word[index] {
				continue
			}
			visited[nx][ny] = true
			index++
			if check(nx, ny) {
				return true
			}
			index--
			visited[nx][ny] = false
		}
		return false
	}
	for i := range board {
		for j := range board[0] {
			if board[i][j] != word[0] {
				continue
			}
			visited[i][j] = true
			index = 1
			if check(i, j) {
				return true
			}
			visited[i][j] = false
		}
	}
	return false
}
