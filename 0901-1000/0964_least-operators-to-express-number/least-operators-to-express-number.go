package main

// Tags:
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
