package main

// Difficulty:
// Medium

// Tags:
// Backtracking

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	o := baseCosts[0]
	var backtrack func(i, cost int)
	backtrack = func(i, cost int) {
		if abs(o-target) > abs(cost-target) || abs(o-target) == abs(cost-target) && o > cost {
			o = cost
		}
		if cost >= target || i == len(toppingCosts) {
			return
		}
		backtrack(i+1, cost)
		backtrack(i+1, cost+toppingCosts[i])
		backtrack(i+1, cost+toppingCosts[i]*2)
	}
	for _, baseCost := range baseCosts {
		backtrack(0, baseCost)
	}
	return o
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
