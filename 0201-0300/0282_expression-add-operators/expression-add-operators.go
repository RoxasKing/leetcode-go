package main

import (
	"strconv"
)

// Tags:
// Backtracking

func addOperators(num string, target int) []string {
	n := len(num)
	arr := make([]int, 0, n)
	for i := range num {
		arr = append(arr, int(num[i]-'0'))
	}
	out := []string{}
	backtrack(arr, n, 0, "", 0, []int{}, target, &out)
	return out
}

func backtrack(arr []int, n, i int, formula string, cur int, nums []int, target int, out *[]string) {
	if i == n {
		res := 0
		for _, num := range nums {
			res += num
		}
		if res == target {
			*out = append(*out, formula)
		}
		return
	}

	preCur := cur
	cur = cur*10 + arr[i]

	if (preCur != 0 || cur != 0) && i < n-1 {
		backtrack(arr, n, i+1, formula, cur, nums, target, out)
	}

	if len(nums) == 0 {
		nums = append(nums, cur)
		backtrack(arr, n, i+1, strconv.Itoa(cur), 0, nums, target, out)
		return
	}

	nums = append(nums, cur)
	backtrack(arr, n, i+1, formula+"+"+strconv.Itoa(cur), 0, nums, target, out)
	nums = nums[:len(nums)-1]

	nums = append(nums, -cur)
	backtrack(arr, n, i+1, formula+"-"+strconv.Itoa(cur), 0, nums, target, out)
	nums = nums[:len(nums)-1]

	pre := nums[len(nums)-1]
	nums[len(nums)-1] = pre * cur
	backtrack(arr, n, i+1, formula+"*"+strconv.Itoa(cur), 0, nums, target, out)
	nums[len(nums)-1] = pre
}
