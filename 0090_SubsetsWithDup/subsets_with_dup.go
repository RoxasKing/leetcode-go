package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	out := [][]int{}
	sort.Ints(nums)

	var dfs func(int, []int)
	dfs = func(index int, tmp []int) {
		out = append(out, tmp)
		n := len(tmp) + 1
		for i := index; i < len(nums); i++ {
			// if 语句的含义是，
			// 在 nums[index:] 中，每个数字只能附着到 temp 中一次
			// 判断方法是 A[i] != A[i-1]
			// 但是 nums[index] == nums[index-1] 也没有关系
			// 因为 nums[index-1] 不在 nums[index:] 中
			// 而且，需要执行 nums[i]!=nums[i-1] 时，
			// 可以肯定 i>=1，所以，不需要验证 i-1>=0
			if i == index || nums[i] != nums[i-1] {
				dfs(i+1, append(tmp, nums[i])[:n:n])
			}
		}
	}
	tmp := make([]int, 0, 0)
	dfs(0, tmp)

	return out
}

func main() {
	nums := []int{1, 2, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
