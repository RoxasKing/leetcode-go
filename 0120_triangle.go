package leetcode

/*
  给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
  例如，给定三角形：
    [
         [2],
        [3,4],
       [6,5,7],
      [4,1,8,3]
    ]
  自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
  说明：
  如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
*/

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += Min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func minimumTotal2(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	save := make([]int, len(triangle))
	copy(save, triangle[len(triangle)-1])
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := range triangle[i] {
			save[j] = triangle[i][j] + Min(save[j], save[j+1])
		}
	}
	return save[0]
}
