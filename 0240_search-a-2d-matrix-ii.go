package main

/*
  编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
    每行的元素从左到右升序排列。
    每列的元素从上到下升序排列。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/search-a-2d-matrix-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func searchMatrixII(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	bSearch := func(array []int) bool {
		l, r := 0, len(array)-1
		for l <= r {
			m := l + (r-l)>>1
			if array[m] < target {
				l = m + 1
			} else if array[m] > target {
				r = m - 1
			} else {
				return true
			}
		}
		return false
	}
	for i := range matrix {
		if matrix[i][len(matrix[i])-1] < target {
			continue
		} else if matrix[i][0] > target {
			return false
		}
		if bSearch(matrix[i]) {
			return true
		}
	}
	return false
}

func searchMatrixII2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, col := len(matrix)-1, 0
	for row >= 0 && col <= len(matrix[0])-1 {
		if matrix[row][col] > target {
			row--
		} else if matrix[row][col] < target {
			col++
		} else {
			return true
		}
	}
	return false
}
