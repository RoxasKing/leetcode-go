package main

import (
	"math"
)

// Tags:
// Backtracking
func judgePoint24(nums []int) bool {
	newNums := make([]float64, len(nums))
	for i := range newNums {
		newNums[i] = float64(nums[i])
	}
	return backtrack(newNums)
}

func backtrack(nums []float64) bool {
	if len(nums) == 1 {
		return math.Abs(nums[0]-24) < 1e-9
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			newNums := make([]float64, 0, len(nums)-1)
			for k := 0; k < len(nums); k++ {
				if k != i && k != j {
					newNums = append(newNums, nums[k])
				}
			}
			for op := 0; op < 4; op++ {
				if i > j && op < 2 || nums[j] == 0 && op == 3 {
					continue
				}
				var res float64
				switch op {
				case 0: // +
					res = nums[i] + nums[j]
				case 1: // *
					res = nums[i] * nums[j]
				case 2: // -
					res = nums[i] - nums[j]
				case 3: // /
					res = nums[i] / nums[j]
				}
				newNums = append(newNums, res)
				if backtrack(newNums) {
					return true
				}
				newNums = newNums[:len(newNums)-1]
			}
		}
	}
	return false
}

// Backtracking
func judgePoint24_2(nums []int) bool {
	newNums := make([]float64, len(nums))
	for i := range newNums {
		newNums[i] = float64(nums[i])
	}
	return backtrack2(newNums, 0)
}

func backtrack2(nums []float64, start int) bool {
	if start == len(nums) {
		return check(nums)
	}
	for i := start; i < len(nums); i++ {
		nums[i], nums[start] = nums[start], nums[i]
		if backtrack2(nums, start+1) {
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
				if res, isValid = cal(exp1, exp2, k); !isValid || math.Abs(res-24) > 1e-9 {
					continue
				}
				return true
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
				if res, isValid = cal(exp2, nums[3], k); !isValid || math.Abs(res-24) > 1e-9 {
					continue
				}
				return true
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
				if res, isValid = cal(exp2, nums[3], k); !isValid || math.Abs(res-24) > 1e-9 {
					continue
				}
				return true
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
				if res, isValid = cal(nums[0], exp2, k); !isValid || math.Abs(res-24) > 1e-9 {
					continue
				}
				return true
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
				if res, isValid = cal(nums[0], exp2, k); !isValid || math.Abs(res-24) > 1e-9 {
					continue
				}
				return true
			}
		}
	}
	return false
}

func cal(a, b float64, op int) (res float64, isValid bool) {
	if a < b && op < 2 || b == 0 && op == 3 { // a+b == b+a and a*b == b*a, so cal once
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
		return a / b, true
	}
	return
}
