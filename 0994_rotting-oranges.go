package main

/*
  在给定的网格中，每个单元格可以有以下三个值之一：
    值 0 代表空单元格；
    值 1 代表新鲜橘子；
    值 2 代表腐烂的橘子。
  每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。
  返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1。
*/

func orangesRotting(grid [][]int) int {
	dr := []int{-1, 0, 1, 0}
	dc := []int{0, -1, 0, 1}
	var queue []int // 用于存放每一批次新增的腐烂橘子
	var remain int  // 剩余新鲜橘子
	var count int   // 分钟数统计
	for i := range grid {
		for j := range grid[i] {
			num := i*len(grid[0]) + j
			if grid[i][j] == 1 {
				remain++
			} else if grid[i][j] == 2 {
				queue = append(queue, num)
			}
		}
	}
	if remain == 0 {
		return 0
	}
	queue = append(queue, -1) // 哨兵，用于监测当前批次的腐烂橘子是否完成感染
	for len(queue) != 0 {
		if queue[0] == -1 { // 若当前批次已经感染完成
			queue = queue[1:]    // 删除当前批次哨兵
			if len(queue) == 0 { // 若已无新增感染橘子，结束循环
				break
			}
			count++                   // 若存在下一批次的橘子，分钟数加1
			queue = append(queue, -1) // 新增下一批次腐烂橘子的哨兵
			continue
		}
		row, col := queue[0]/len(grid[0]), queue[0]%len(grid[0])
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			nr, nc := row+dr[i], col+dc[i]
			num := nr*len(grid[0]) + nc
			if 0 <= nr && nr < len(grid) && 0 <= nc && nc < len(grid[0]) && grid[nr][nc] == 1 {
				queue = append(queue, num)
				grid[nr][nc] = 2
				remain--
			}
		}
	}
	if remain == 0 {
		return count
	}
	return -1
}
