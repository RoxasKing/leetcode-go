package main

import (
	"strconv"
)

/*
  Given a string num that contains only digits and an integer target, return all possibilities to add the binary operators '+', '-', or '*' between the digits of num so that the resultant expression evaluates to the target value.

  Example 1:
    Input: num = "123", target = 6
    Output: ["1*2*3","1+2+3"]

  Example 2:
    Input: num = "232", target = 8
    Output: ["2*3+2","2+3*2"]

  Example 3:
    Input: num = "105", target = 5
    Output: ["1*0+5","10-5"]

  Example 4:
    Input: num = "00", target = 0
    Output: ["0*0","0+0","0-0"]

  Example 5:
    Input: num = "3456237490", target = 9191
    Output: []

  Constraints:
    1. 1 <= num.length <= 10
    2. num consists of only digits.
    3. -2^31 <= target <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/expression-add-operators
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

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
