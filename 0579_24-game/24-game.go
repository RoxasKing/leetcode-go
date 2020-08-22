package main

import "math"

/*
  你有 4 张写有 1 到 9 数字的牌。你需要判断是否能通过 *，/，+，-，(，) 的运算得到 24。

  注意:
    除法运算符 / 表示实数除法，而不是整数除法。例如 4 / (1 - 2/3) = 12 。
    每个运算符对两个数进行运算。特别是我们不能用 - 作为一元运算符。例如，[1, 1, 1, 1] 作为输入时，表达式 -1 - 1 - 1 - 1 是不允许的。
    你不能将数字连接在一起。例如，输入为 [1, 2, 1, 2] 时，不能写成 12 + 12 。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/24-game
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func judgePoint24(nums []int) bool {
	return backTrack(nums, 0)
}

func backTrack(nums []int, start int) bool {
	if start == len(nums) {
		return check(nums)
	}
	for i := start; i < len(nums); i++ {
		nums[i], nums[start] = nums[start], nums[i]
		if backTrack(nums, start+1) {
			return true
		}
		nums[i], nums[start] = nums[start], nums[i]
	}
	return false
}

var ops = []byte{'+', '-', '*', '/'}

func check(nums []int) bool {
	var (
		res        float64
		isValid    bool
		exp1, exp2 float64
	)

	// (a op1 b) op3 (c op2 d)
	for i := range ops {
		if exp1, isValid = cal(float64(nums[0]), float64(nums[1]), ops[i]); !isValid {
			continue
		}
		for j := range ops {
			if exp2, isValid = cal(float64(nums[2]), float64(nums[3]), ops[j]); !isValid {
				continue
			}
			for k := range ops {
				if res, isValid = cal(exp1, exp2, ops[k]); !isValid {
					continue
				}
				if math.Abs(res-24) <= 1e-9 {
					return true
				}
			}
		}
	}

	// ((a op1 b) op2 c) op3 d
	for i := range ops {
		if exp1, isValid = cal(float64(nums[0]), float64(nums[1]), ops[i]); !isValid {
			continue
		}
		for j := range ops {
			if exp2, isValid = cal(exp1, float64(nums[2]), ops[j]); !isValid {
				continue
			}
			for k := range ops {
				if res, isValid = cal(exp2, float64(nums[3]), ops[k]); !isValid {
					continue
				}
				if math.Abs(res-24) <= 1e-9 {
					return true
				}
			}
		}
	}

	// (a op1 (b op2 c)) op3 d
	for i := range ops {
		if exp1, isValid = cal(float64(nums[1]), float64(nums[2]), ops[i]); !isValid {
			continue
		}
		for j := range ops {
			if exp2, isValid = cal(float64(nums[0]), exp1, ops[j]); !isValid {
				continue
			}
			for k := range ops {
				if res, isValid = cal(exp2, float64(nums[3]), ops[k]); !isValid {
					continue
				}
				if math.Abs(res-24) <= 1e-9 {
					return true
				}
			}
		}
	}

	// a op1 ((b op2 c) op3 d)
	for i := range ops {
		if exp1, isValid = cal(float64(nums[1]), float64(nums[2]), ops[i]); !isValid {
			continue
		}
		for j := range ops {
			if exp2, isValid = cal(exp1, float64(nums[3]), ops[j]); !isValid {
				continue
			}
			for k := range ops {
				if res, isValid = cal(float64(nums[0]), exp2, ops[k]); !isValid {
					continue
				}
				if math.Abs(res-24) <= 1e-9 {
					return true
				}
			}
		}
	}

	// a op1 (b op2 (c op3 d))
	for i := range ops {
		if exp1, isValid = cal(float64(nums[2]), float64(nums[3]), ops[i]); !isValid {
			continue
		}
		for j := range ops {
			if exp2, isValid = cal(float64(nums[1]), exp1, ops[j]); !isValid {
				continue
			}
			for k := range ops {
				if res, isValid = cal(float64(nums[0]), exp2, ops[k]); !isValid {
					continue
				}
				if math.Abs(res-24) <= 1e-9 {
					return true
				}
			}
		}
	}
	return false
}

func cal(a, b float64, op byte) (res float64, isValid bool) {
	switch op {
	case '+':
		return a + b, true
	case '-':
		return a - b, true
	case '*':
		return a * b, true
	case '/':
		if b == 0 {
			return 0, false
		}
		return a / b, true
	}
	return
}
