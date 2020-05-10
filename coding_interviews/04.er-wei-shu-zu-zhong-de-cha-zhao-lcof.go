package codinginterviews

/*
  在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	headRow, tailCol := 0, len(matrix[0])-1
	for headRow < len(matrix) && tailCol >= 0 {
		if matrix[headRow][tailCol] < target {
			headRow++
		} else if matrix[headRow][tailCol] > target {
			tailCol--
		} else {
			return true
		}
	}
	return false
}
