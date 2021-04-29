package main

/*
  Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0). Find two lines, which, together with the x-axis forms a container, such that the container contains the most water.

  Notice that you may not slant the container.

  Example 1:
    Input: height = [1,8,6,2,5,4,8,3,7]
    Output: 49
    Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

  Example 2:
    Input: height = [1,1]
    Output: 1

  Example 3:
    Input: height = [4,3,2,1,4]
    Output: 16

  Example 4:
    Input: height = [1,2,1]
    Output: 2

  Constraints:
    1. n == height.length
    2. 2 <= n <= 10^5
    3. 0 <= height[i] <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/container-with-most-water
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Two Pointers
func maxArea(height []int) int {
	out := 0
	for l, r := 0, len(height)-1; l < r; {
		out = Max(out, Min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return out
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
