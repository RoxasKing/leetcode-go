package main

/*
  给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

  示例 1：
    输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
    输出：6
    解释：最大矩形如上图所示。

  示例 2：
    输入：matrix = []
    输出：0

  示例 3：
    输入：matrix = [["0"]]
    输出：0

  示例 4：
    输入：matrix = [["1"]]
    输出：1

  示例 5：
    输入：matrix = [["0","0"]]
    输出：0

  提示：
    rows == matrix.length
    cols == matrix.length
    0 <= row, cols <= 200
    matrix[i][j] 为 '0' 或 '1'

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximal-rectangle
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Stack
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	var out int
	m, n := len(matrix), len(matrix[0])
	heights := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		out = Max(out, largestRectangleArea(heights))
	}
	return out
}

func largestRectangleArea(heights []int) int {
	out := 0
	stack := MakeStack()
	stack.Push(-1)
	for i := range heights {
		for stack.Top() != -1 && heights[i] <= heights[stack.Top()] {
			out = Max(out, heights[stack.Pop()]*(i-1-stack.Top()))
		}
		stack.Push(i)
	}
	for stack.Top() != -1 {
		out = Max(out, heights[stack.Pop()]*(len(heights)-1-stack.Top()))
	}
	return out
}

type Stack []int

func MakeStack() Stack {
	return Stack{}
}

func (s Stack) Size() int {
	return len(s)
}

func (s Stack) Top() int {
	last := len(s) - 1
	return s[last]
}

func (s *Stack) Pop() int {
	last := len(*s) - 1
	out := (*s)[last]
	*s = (*s)[:last]
	return out
}

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
