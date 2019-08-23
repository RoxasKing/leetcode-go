package main

import "fmt"

func subsets(nums []int) [][]int {
	out := make([][]int, 0)

	var dfs func(int, int, []int)
	dfs = func(first, size int, curr []int) {
		if len(curr) == size {
			temp := make([]int, size)
			copy(temp, curr)
			out = append(out, temp)
			return
		}

		for i := first; i < len(nums)-(size-len(curr))+1; i++ {
			curr = append(curr, nums[i])
			dfs(i+1, size, curr)
			curr = curr[:len(curr)-1]
		}
	}
	for i := 0; i <= len(nums); i++ {
		dfs(0, i, make([]int, 0))
	}

	return out
}

func main() {
	nums := []int{1, 2, 4}
	fmt.Println(subsets(nums))
}
