package main

/*
  「以扣会友」线下活动所在场地由若干主题空间与走廊组成，场地的地图记作由一维字符串型数组 grid，字符串中仅包含 "0"～"5" 这 6 个字符。地图上每一个字符代表面积为 1 的区域，其中 "0" 表示走廊，其他字符表示主题空间。相同且连续（连续指上、下、左、右四个方向连接）的字符组成同一个主题空间。

  假如整个 grid 区域的外侧均为走廊。请问，不与走廊直接相邻的主题空间的最大面积是多少？如果不存在这样的空间请返回 0。

  示例 1:
    输入：grid = ["110","231","221"]
    输出：1
    解释：4 个主题空间中，只有 1 个不与走廊相邻，面积为 1。

  示例 2:
    输入：grid = ["11111100000","21243101111","21224101221","11111101111"]
    输出：3
    解释：8 个主题空间中，有 5 个不与走廊相邻，面积分别为 3、1、1、1、2，最大面积为 3。

  提示：
    1. 1 <= grid.length <= 500
    2. 1 <= grid[i].length <= 500
    3. grid[i][j] 仅可能是 "0"～"5"

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/YesdPw
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS

func largestArea(grid []string) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				visited[i][j] = true
				if i+1 < m && grid[i+1][j] != '0' {
					walk(grid, visited, m, n, i+1, j, grid[i+1][j])
				}
				if i-1 >= 0 && grid[i-1][j] != '0' {
					walk(grid, visited, m, n, i-1, j, grid[i-1][j])
				}
				if j+1 < n && grid[i][j+1] != '0' {
					walk(grid, visited, m, n, i, j+1, grid[i][j+1])
				}
				if j-1 >= 0 && grid[i][j-1] != '0' {
					walk(grid, visited, m, n, i, j-1, grid[i][j-1])
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		if !visited[i][0] {
			walk(grid, visited, m, n, i, 0, grid[i][0])
		}
		if !visited[i][n-1] {
			walk(grid, visited, m, n, i, n-1, grid[i][n-1])
		}
	}
	for j := 0; j < n; j++ {
		if !visited[0][j] {
			walk(grid, visited, m, n, 0, j, grid[0][j])
		}
		if !visited[m-1][j] {
			walk(grid, visited, m, n, m-1, j, grid[m-1][j])
		}
	}

	out := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] {
				continue
			}
			q := [][]int{{i, j}}
			visited[i][j] = true
			cnt := 0
			for len(q) > 0 {
				x, y := q[0][0], q[0][1]
				q = q[1:]
				cnt++
				if x+1 < m && !visited[x+1][y] && grid[x+1][y] == grid[x][y] {
					visited[x+1][y] = true
					q = append(q, []int{x + 1, y})
				}
				if x-1 >= 0 && !visited[x-1][y] && grid[x-1][y] == grid[x][y] {
					visited[x-1][y] = true
					q = append(q, []int{x - 1, y})
				}
				if y+1 < n && !visited[x][y+1] && grid[x][y+1] == grid[x][y] {
					visited[x][y+1] = true
					q = append(q, []int{x, y + 1})
				}
				if y-1 >= 0 && !visited[x][y-1] && grid[x][y-1] == grid[x][y] {
					visited[x][y-1] = true
					q = append(q, []int{x, y - 1})
				}
			}
			out = Max(out, cnt)
		}
	}
	return out
}

func walk(grid []string, visited [][]bool, m, n, i, j int, val byte) {
	if i < 0 || i > m-1 || j < 0 || j > n-1 || visited[i][j] || grid[i][j] != val {
		return
	}
	visited[i][j] = true
	walk(grid, visited, m, n, i+1, j, val)
	walk(grid, visited, m, n, i, j+1, val)
	walk(grid, visited, m, n, i-1, j, val)
	walk(grid, visited, m, n, i, j-1, val)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
