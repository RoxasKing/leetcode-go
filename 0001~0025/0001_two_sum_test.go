package _001_0025

import (
	"fmt"
	"testing"
)

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
