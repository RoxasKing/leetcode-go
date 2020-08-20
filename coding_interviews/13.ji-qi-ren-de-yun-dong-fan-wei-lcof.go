package codinginterviews

/*
  地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
*/

// BFS + Recursion
func movingCount(m int, n int, k int) int {
	steps := [][]int{{1, 0}, {0, 1}}
	isValid := func(x, y int) bool {
		var sum int
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
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}
	used[0][0] = true
	var max int
	count := 1
	var bfs func(int, int)
	bfs = func(x, y int) {
		max = Max(max, count)
		for _, s := range steps {
			nx, ny := x+s[0], y+s[1]
			if nx < 0 || nx > m-1 || ny < 0 || ny > n-1 || used[nx][ny] || !isValid(nx, ny) {
				continue
			}
			count++
			used[nx][ny] = true
			bfs(nx, ny)
		}
	}
	bfs(0, 0)
	return max
}

// BFS + Queue
func movingCount2(m int, n int, k int) int {
	steps := [][]int{{1, 0}, {0, 1}}
	isValid := func(x, y int) bool {
		var sum int
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
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}
	used[0][0] = true
	count := 1
	queue := [][]int{{0, 0}}
	for len(queue) != 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		for _, s := range steps {
			nx, ny := x+s[0], y+s[1]
			if nx < 0 || nx > m-1 || ny < 0 || ny > n-1 || used[nx][ny] || !isValid(nx, ny) {
				continue
			}
			count++
			used[nx][ny] = true
			queue = append(queue, []int{nx, ny})
		}
	}
	return count
}

//
func movingCount3(m int, n int, k int) int {
	isValid := func(x, y int) bool {
		var sum int
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
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}
	used[0][0] = true
	count := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isValid(i, j) && (i-1 >= 0 && used[i-1][j] || j-1 >= 0 && used[i][j-1]) {
				used[i][j] = true
				count++
			}
		}
	}
	return count
}
