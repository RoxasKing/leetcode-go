package main

/*
  输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

  示例 1：
    输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
    输出：[1,2,3,6,9,8,7,4,5]

  示例 2：
    输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
    输出：[1,2,3,4,8,12,11,10,9,5,6,7]

  限制：
    0 <= matrix.length <= 100
    0 <= matrix[i].length <= 100

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	l, r, u, d := 0, n-1, 0, m-1
	out := make([]int, 0, m*n)
	f := 0
	for l <= r && u <= d {
		switch f {
		case 0:
			for j := l; j <= r; j++ {
				out = append(out, matrix[u][j])
			}
			u++
		case 1:
			for i := u; i <= d; i++ {
				out = append(out, matrix[i][r])
			}
			r--
		case 2:
			for j := r; j >= l; j-- {
				out = append(out, matrix[d][j])
			}
			d--
		case 3:
			for i := d; i >= u; i-- {
				out = append(out, matrix[i][l])
			}
			l++
		}
		f = (f + 1) % 4
	}
	return out
}
