package main

/*
  地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

  示例 1：
    输入：m = 2, n = 3, k = 1
    输出：3

  示例 2：
    输入：m = 3, n = 1, k = 0
    输出：1

  提示：
    1. 1 <= n,m <= 100
    2. 0 <= k <= 20

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// BFS + Queue
func movingCount(m int, n int, k int) int {
	out := 0
	queue := [][]int{{0, 0}}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		x, y := point[0], point[1]
		if x < 0 || m-1 < x || y < 0 || n-1 < y || visited[x][y] {
			continue
		}
		visited[x][y] = true
		if isValid(x, y, k) {
			out++
			queue = append(queue, []int{x + 1, y}, []int{x, y + 1})
		}
	}
	return out
}

func isValid(x, y, k int) bool {
	sum := 0
	for x != 0 {
		sum += x % 10
		x /= 10
	}
	for y != 0 {
		sum += y % 10
		y /= 10
	}
	return sum <= k
}

func movingCount2(m int, n int, k int) int {
	out := 1
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	visited[0][0] = true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i > 0 && visited[i-1][j] || j > 0 && visited[i][j-1]) && isValid(i, j, k) {
				out++
				visited[i][j] = true
			}
		}
	}
	return out
}
