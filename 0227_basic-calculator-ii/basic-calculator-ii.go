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
	l, r := 0, 0
	for l < len(s) {
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
			lastNumIdx := len(nums) - 1
			num1, num2 := nums[lastNumIdx-1], nums[lastNumIdx]
			nums = nums[:lastNumIdx-1]
			lastOpsIdx := len(ops) - 1
			if ops[lastOpsIdx] == '*' {
				nums = append(nums, num1*num2)
			} else if ops[lastOpsIdx] == '/' {
				nums = append(nums, num1/num2)
			}
			ops = ops[:lastOpsIdx]
		}

		for r < len(s) && s[r] == ' ' {
			r++
		}

		if r < len(s) {
			ops = append(ops, s[r])
		}

		l = r + 1
	}
	out := nums[0]
	nums = nums[1:]
	for len(nums) != 0 {
		if ops[0] == '+' {
			out += nums[0]
		} else if ops[0] == '-' {
			out -= nums[0]
		}
		nums = nums[1:]
		ops = ops[1:]
	}
	return out
}
