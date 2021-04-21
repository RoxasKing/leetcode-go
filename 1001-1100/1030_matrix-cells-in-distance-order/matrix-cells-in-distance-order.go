package main

import "sort"

/*
  给出 R 行 C 列的矩阵，其中的单元格的整数坐标为 (r, c)，满足 0 <= r < R 且 0 <= c < C。
  另外，我们在该矩阵中给出了一个坐标为 (r0, c0) 的单元格。
  返回矩阵中的所有单元格的坐标，并按到 (r0, c0) 的距离从最小到最大的顺序排，其中，两单元格(r1, c1) 和 (r2, c2) 之间的距离是曼哈顿距离，|r1 - r2| + |c1 - c2|。（你可以按任何满足此条件的顺序返回答案。）

  示例 1：
    输入：R = 1, C = 2, r0 = 0, c0 = 0
    输出：[[0,0],[0,1]]
    解释：从 (r0, c0) 到其他单元格的距离为：[0,1]

  示例 2：
    输入：R = 2, C = 2, r0 = 0, c0 = 1
    输出：[[0,1],[0,0],[1,1],[1,0]]
    解释：从 (r0, c0) 到其他单元格的距离为：[0,1,1,2]
    [[0,1],[1,1],[0,0],[1,0]] 也会被视作正确答案。

  示例 3：
    输入：R = 2, C = 3, r0 = 1, c0 = 2
    输出：[[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]
    解释：从 (r0, c0) 到其他单元格的距离为：[0,1,1,2,2,3]
    其他满足题目要求的答案也会被视为正确，例如 [[1,2],[1,1],[0,2],[1,0],[0,1],[0,0]]。

  提示：
    1. 1 <= R <= 100
    2. 1 <= C <= 100
    3. 0 <= r0 < R
    4. 0 <= c0 < C

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/matrix-cells-in-distance-order
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Bucket Sort
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	maxDist := Max(r0, R-1-r0) + Max(c0, C-1-c0)
	buckets := make([][][]int, maxDist+1)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			dist := Abs(i-r0) + Abs(j-c0)
			buckets[dist] = append(buckets[dist], []int{i, j})
		}
	}
	out := make([][]int, 0, R*C)
	for _, b := range buckets {
		out = append(out, b...)
	}
	return out
}

// Sort
func allCellsDistOrder2(R int, C int, r0 int, c0 int) [][]int {
	out := make([][]int, 0, R*C)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			out = append(out, []int{i, j})
		}
	}
	sort.Slice(out, func(i, j int) bool {
		return Abs(out[i][0]-r0)+Abs(out[i][1]-c0) < Abs(out[j][0]-r0)+Abs(out[j][1]-c0)
	})
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
