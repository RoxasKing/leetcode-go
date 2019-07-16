package main

import "fmt"

func bigger(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func smaller(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func trap(height []int) int {
	length := len(height)
	if length <= 2 {
		return 0
	}

	left, right := make([]int, length), make([]int, length)

	left[0], right[length-1] = height[0], height[length-1]

	for i := 1; i < length; i++ {
		left[i] = bigger(left[i-1], height[i])
		right[length-1-i] = bigger(right[length-i], height[length-1-i])

		// left[i] 是 height[:i+1] 中的最大值
		// 相当于从左到右的投影
		// right[i] 是 height[i:] 中的最大值
		// 相当于从右到左的投影
	}

	water := 0
	for i := 0; i < length; i++ {
		// 存水量取决于 左右最大值 中的较小值
		// 相当于取两个投影的交集，然后用交集去减相应的原数组
		water += smaller(left[i], right[i]) - height[i]
	}

	return water
}

func main() {
	in := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	//in := []int{2, 0, 2}
	fmt.Println(trap(in))
}
