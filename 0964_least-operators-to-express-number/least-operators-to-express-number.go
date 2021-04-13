package main

/*
  Given a single positive integer x, we will write an expression of the form x (op1) x (op2) x (op3) x ... where each operator op1, op2, etc. is either addition, subtraction, multiplication, or division (+, -, *, or /). For example, with x = 3, we might write 3 * 3 / 3 + 3 - 3 which is a value of 3.

  When writing such an expression, we adhere to the following conventions:
    1. The division operator (/) returns rational numbers.
    2. There are no parentheses placed anywhere.
    3. We use the usual order of operations: multiplication and division happen before addition and subtraction.
    4. It is not allowed to use the unary negation operator (-). For example, "x - x" is a valid expression as it only uses subtraction, but "-x + x" is not because it uses negation.
  We would like to write an expression with the least number of operators such that the expression equals the given target. Return the least number of operators used.

  Example 1:
    Input: x = 3, target = 19
    Output: 5
    Explanation: 3 * 3 + 3 * 3 + 3 / 3.
    The expression contains 5 operations.

  Example 2:
    Input: x = 5, target = 501
    Output: 8
    Explanation: 5 * 5 * 5 * 5 - 5 * 5 * 5 + 5 / 5.
    The expression contains 8 operations.

  Example 3:
    Input: x = 100, target = 100000000
    Output: 3
    Explanation: 100 * 100 * 100 * 100.
    The expression contains 3 operations.

  Constraints:
    2 <= x <= 100
    1 <= target <= 2 * 108

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/least-operators-to-express-number
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS + Dynamic Programming
func leastOpsExpressTarget(x int, target int) int {
	memo := make(map[int]int)
	dfs(x, target, memo)
	return memo[target]
}

func dfs(x int, target int, memo map[int]int) {
	if _, ok := memo[target]; ok {
		return
	}

	if target < x {
		memo[target] = Min(target*2-1, (x-target)*2)
		return
	}

	out := target*2 - 1
	i := 1
	product := x
	for product < target {
		tmp := (target/product)*i - 1
		remain := target % product
		if remain > 0 {
			dfs(x, remain, memo)
			tmp += 1 + memo[remain]
		}
		out = Min(out, tmp)
		product *= x
		i++
	}

	if product == target {
		memo[target] = i - 1
		return
	}

	if product-target > target {
		memo[target] = out
		return
	}

	dfs(x, product-target, memo)
	out = Min(out, i+memo[product-target])
	memo[target] = out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
