package main

// Tags:
// Backtracking
func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	out := baseCosts[0]
	for _, bc := range baseCosts {
		dfs(toppingCosts, 0, bc, target, &out)
	}
	return out
}

func dfs(toppingCosts []int, i int, cost, target int, out *int) {
	if Abs(cost-target) < Abs(*out-target) {
		*out = cost
	} else if Abs(cost-target) == Abs(*out-target) {
		*out = Min(*out, cost)
	}

	if cost >= target || i == len(toppingCosts) {
		return
	}

	for t := 0; t <= 2; t++ {
		dfs(toppingCosts, i+1, cost+toppingCosts[i]*t, target, out)
	}
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
