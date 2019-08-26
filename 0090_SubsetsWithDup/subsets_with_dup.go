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
	nums := []int{1, 2, 2}
	aa := nums[:3:3]
	fmt.Println(subsetsWithDup(nums))
	fmt.Println(aa)
}
