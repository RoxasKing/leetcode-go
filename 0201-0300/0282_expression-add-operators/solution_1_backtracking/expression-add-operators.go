package main

import "strconv"

// Difficulty:
// Hard

// Tags:
// Backtracking

func addOperators(num string, target int) []string {
	out := []string{}
	dfs(num, "", 0, 0, target, []int{}, &out)
	return out
}

func dfs(num, exp string, i, v, target int, nums []int, out *[]string) {
	if len(num) == i {
		res := 0
		for _, num := range nums {
			res += num
		}
		if res == target {
			*out = append(*out, exp)
		}
		return
	}

	pv := v
	v = v*10 + int(num[i]-'0')
	if (pv != 0 || v != 0) && i < len(num)-1 {
		dfs(num, exp, i+1, v, target, nums, out)
	}

	if len(nums) == 0 {
		dfs(num, strconv.Itoa(v), i+1, 0, target, []int{v}, out)
		return
	}

	t := nums[len(nums)-1]
	nums[len(nums)-1] *= v
	dfs(num, exp+"*"+strconv.Itoa(v), i+1, 0, target, nums, out)
	nums[len(nums)-1] = t
	nums = append(nums, v)
	dfs(num, exp+"+"+strconv.Itoa(v), i+1, 0, target, nums, out)
	nums[len(nums)-1] = -v
	dfs(num, exp+"-"+strconv.Itoa(v), i+1, 0, target, nums, out)
}
