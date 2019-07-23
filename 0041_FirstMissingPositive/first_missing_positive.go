package main

import "fmt"

func firstMissingPositive(nums []int) int {
	size := len(nums)
	dict := make(map[int]bool)
	// 存储所有数字
	for i := 0; i < size; i++ {
		dict[nums[i]] = true
	}
	// 默认最小正整数为 1
	out := 1
	for {
		if _, ok := dict[out]; !ok {
			// 如果当前最小正整数没有在字典中出现过, 就输出
			return out
		}
		out++
	}
}

func main() {
	//in := []int{1, 2, 0}
	//in := []int{3, 4, -1, 1}
	//in := []int{7, 8, 9, 11, 12}
	//in := []int{-1}
	in := []int{3, 4, -1, 1}
	fmt.Println(firstMissingPositive(in))
}
