package main

/*
  请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

    [["a","b","c","e"],
    ["s","f","c","s"],
    ["a","d","e","e"]]
  但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。

  示例 1：
    输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
    输出：true

  示例 2：
    输入：board = [["a","b"],["c","d"]], word = "abcd"
    输出：false

  提示：
    1. 1 <= board.length <= 200
    2. 1 <= board[i].length <= 200

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, word, visited, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, visited [][]bool, i, x, y int) bool {
	if i == len(word) {
		return true
	}
	if x < 0 || x > len(board)-1 || y < 0 || y > len(board[0])-1 || visited[x][y] || board[x][y] != word[i] {
		return false
	}
	visited[x][y] = true
	out := dfs(board, word, visited, i+1, x-1, y) ||
		dfs(board, word, visited, i+1, x+1, y) ||
		dfs(board, word, visited, i+1, x, y-1) ||
		dfs(board, word, visited, i+1, x, y+1)
	visited[x][y] = false
	return out
}
