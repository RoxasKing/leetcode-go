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

// Backtracking
func judgePoint24(nums []int) bool {
	newNums := make([]float64, len(nums))
	for i := range newNums {
		newNums[i] = float64(nums[i])
	}
	return backTrack(newNums, 0)
}

func backTrack(nums []float64, start int) bool {
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

func check(nums []float64) bool {
	var (
		res        float64
		isValid    bool
		exp1, exp2 float64
	)

	// (a op1 b) op3 (c op2 d)
	for i := 0; i < 4; i++ {
		if exp1, isValid = cal(nums[0], nums[1], i); !isValid {
			continue
		}
		for j := 0; j < 4; j++ {
			if exp2, isValid = cal(nums[2], nums[3], j); !isValid {
				continue
			}
			for k := 0; k < 4; k++ {
				if res, isValid = cal(exp1, exp2, k); !isValid {
					continue
				}
				if math.Abs(res-24) <= 1e-9 {
					return true
				}
			}
		}
	}

	// ((a op1 b) op2 c) op3 d
	for i := 0; i < 4; i++ {
		if exp1, isValid = cal(nums[0], nums[1], i); !isValid {
			continue
		}
		for j := 0; j < 4; j++ {
			if exp2, isValid = cal(exp1, nums[2], j); !isValid {
				continue
			}
			for k := 0; k < 4; k++ {
				if res, isValid = cal(exp2, nums[3], k); !isValid {
					continue
				}
				if math.Abs(res-24) <= 1e-9 {
					return true
				}
			}
		}
	}

	// (a op1 (b op2 c)) op3 d
	for i := 0; i < 4; i++ {
		if exp1, isValid = cal(nums[1], nums[2], i); !isValid {
			continue
		}
		for j := 0; j < 4; j++ {
			if exp2, isValid = cal(nums[0], exp1, j); !isValid {
				continue
			}
			for k := 0; k < 4; k++ {
				if res, isValid = cal(exp2, nums[3], k); !isValid {
					continue
				}
				if math.Abs(res-24) <= 1e-9 {
					return true
				}
			}
		}
	}

	// a op1 ((b op2 c) op3 d)
	for i := 0; i < 4; i++ {
		if exp1, isValid = cal(nums[1], nums[2], i); !isValid {
			continue
		}
		for j := 0; j < 4; j++ {
			if exp2, isValid = cal(exp1, nums[3], j); !isValid {
				continue
			}
			for k := 0; k < 4; k++ {
				if res, isValid = cal(nums[0], exp2, k); !isValid {
					continue
				}
				if math.Abs(res-24) <= 1e-9 {
					return true
				}
			}
		}
	}

	// a op1 (b op2 (c op3 d))
	for i := 0; i < 4; i++ {
		if exp1, isValid = cal(nums[2], nums[3], i); !isValid {
			continue
		}
		for j := 0; j < 4; j++ {
			if exp2, isValid = cal(nums[1], exp1, j); !isValid {
				continue
			}
			for k := 0; k < 4; k++ {
				if res, isValid = cal(nums[0], exp2, k); !isValid {
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

func cal(a, b float64, op int) (res float64, isValid bool) {
	if a < b && op < 2 { // a+b == b+a and a*b == b*a, so cal once
		return 0, false
	}
	switch op {
	case 0: // +
		return a + b, true
	case 1: // *
		return a * b, true
	case 2: // -
		return a - b, true
	case 3: // /
		if b == 0 {
			return 0, false
		}
		return a / b, true
	}
	return
}
