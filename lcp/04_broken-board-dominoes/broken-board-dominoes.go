package main

/*
  你有一块棋盘，棋盘上有一些格子已经坏掉了。你还有无穷块大小为1 * 2的多米诺骨牌，你想把这些骨牌不重叠地覆盖在完好的格子上，请找出你最多能在棋盘上放多少块骨牌？这些骨牌可以横着或者竖着放。

  输入：n, m代表棋盘的大小；broken是一个b * 2的二维数组，其中每个元素代表棋盘上每一个坏掉的格子的位置。

  输出：一个整数，代表最多能在棋盘上放的骨牌数。

  示例 1：
    输入：n = 2, m = 3, broken = [[1, 0], [1, 1]]
    输出：2
    解释：我们最多可以放两块骨牌：[[0, 0], [0, 1]]以及[[0, 2], [1, 2]]。（见下图）

  示例 2：
    输入：n = 3, m = 3, broken = []
    输出：4
    解释：下图是其中一种可行的摆放方式

  限制：
    1. 1 <= n <= 8
    2. 1 <= m <= 8
    3. 0 <= b <= n * m

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/broken-board-dominoes
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Hungarian Algorithm
func domino(n int, m int, broken [][]int) int {
	inValid := make([]bool, n*m)
	for _, b := range broken {
		i, j := b[0], b[1]
		idx := i*m + j
		inValid[idx] = true
	}

	edges := make([][]int, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			idx := i*m + j
			if (i+j)&1 == 1 || inValid[idx] {
				continue
			}
			if i-1 >= 0 && !inValid[(i-1)*m+j] {
				edges[idx] = append(edges[idx], (i-1)*m+j)
			}
			if i+1 < n && !inValid[(i+1)*m+j] {
				edges[idx] = append(edges[idx], (i+1)*m+j)
			}
			if j-1 >= 0 && !inValid[i*m+j-1] {
				edges[idx] = append(edges[idx], i*m+j-1)
			}
			if j+1 < m && !inValid[i*m+j+1] {
				edges[idx] = append(edges[idx], i*m+j+1)
			}
		}
	}

	match := make([]int, n*m)
	for i := range match {
		match[i] = -1
	}

	out := 0
	for i := range edges {
		visited := make(map[int]bool)
		if find(visited, edges, match, i) {
			out++
		}
	}
	return out
}

func find(visited map[int]bool, edges [][]int, match []int, i int) bool {
	for _, j := range edges[i] {
		if visited[j] {
			continue
		}
		visited[j] = true
		if match[j] == -1 || find(visited, edges, match, match[j]) {
			match[j] = i
			return true
		}
	}
	return false
}
