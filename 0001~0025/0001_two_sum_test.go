package _001_0025

import (
	"fmt"
	"testing"
)

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))

	for i, b := range nums {
		if j, ok := index[target-b]; ok {
			return []int{j, i}
		}
		index[b] = i
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	target := 7
	fmt.Println(twoSum(nums, target))
}
