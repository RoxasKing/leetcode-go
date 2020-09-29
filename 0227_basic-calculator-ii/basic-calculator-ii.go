package main

import "strconv"

/*
  实现一个基本的计算器来计算一个简单的字符串表达式的值。
  字符串表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。

  说明：
    你可以假设所给定的表达式都是有效的。
    请不要使用内置的库函数 eval。
*/

func calculate(s string) int {
	var nums []int
	var ops []byte
	for l, r := 0, 0; l < len(s); l = r + 1 {
		for l < len(s) && s[l] == ' ' {
			l++
		}

		r = l + 1
		for r < len(s) && '0' <= s[r] && s[r] <= '9' {
			r++
		}
		num, _ := strconv.Atoi(s[l:r])
		nums = append(nums, num)

		if len(nums) >= 2 && len(ops) != 0 && (ops[len(ops)-1] == '*' || ops[len(ops)-1] == '/') {
			lastNumIdx, lastOpsIdx := len(nums)-1, len(ops)-1
			num1, num2, op := nums[lastNumIdx-1], nums[lastNumIdx], ops[lastOpsIdx]
			nums, ops = nums[:lastNumIdx-1], ops[:lastOpsIdx]
			if op == '*' {
				nums = append(nums, num1*num2)
			} else {
				nums = append(nums, num1/num2)
			}
		}

		for r < len(s) && s[r] == ' ' {
			r++
		}

		if r < len(s) {
			ops = append(ops, s[r])
		}
	}
	out := nums[0]
	nums = nums[1:]
	for len(nums) != 0 {
		if ops[0] == '+' {
			out += nums[0]
		} else {
			out -= nums[0]
		}
		nums, ops = nums[1:], ops[1:]
	}
	return out
}
