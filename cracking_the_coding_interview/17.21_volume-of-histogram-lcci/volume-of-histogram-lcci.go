package main

/*
  Imagine a histogram (bar graph). Design an algorithm to compute the volume of water it could hold if someone poured water across the top. You can assume that each histogram bar has width 1.

  The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of water (blue section) are being trapped. Thanks Marcos for contributing this image!

  Example:
    Input: [0,1,0,2,1,0,1,3,2,1,2,1]
    Output: 6

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/volume-of-histogram-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Monotone Stack
func trap(height []int) int {
	stack := []int{}
	out := 0
	for i := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			bottom := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			out += (i - 1 - stack[len(stack)-1]) * (Min(height[stack[len(stack)-1]], height[i]) - bottom)
		}
		stack = append(stack, i)
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
