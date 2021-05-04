package main

/*
  Given an array of integers heights representing the histogram's bar height where the width of each bar is 1, return the area of the largest rectangle in the histogram.

  Example 1:
    Input: heights = [2,1,5,6,2,3]
    Output: 10
    Explanation: The above is a histogram where width of each bar is 1.
      The largest rectangle is shown in the red area, which has an area = 10 units.

  Example 2:
    Input: heights = [2,4]
    Output: 4

  Constraints:
    1. 1 <= heights.length <= 10^5
    2. 0 <= heights[i] <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/largest-rectangle-in-histogram
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Monotone Stack
func largestRectangleArea(heights []int) int {
	out := 0
	s := IntStack{-1}
	for i := range heights {
		for s.Len() > 1 && heights[s.Top()] >= heights[i] {
			height := heights[s.Pop()]
			out = Max(out, height*(i-1-s.Top()))
		}
		s.Push(i)
	}
	n := len(heights)
	for s.Len() > 1 {
		height := heights[s.Pop()]
		out = Max(out, height*(n-1-s.Top()))
	}
	return out
}

type IntStack []int

func (s IntStack) Len() int    { return len(s) }
func (s IntStack) Top() int    { return s[s.Len()-1] }
func (s *IntStack) Push(x int) { *s = append(*s, x) }
func (s *IntStack) Pop() int {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
