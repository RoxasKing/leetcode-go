package main

// Difficulty:
// Medium

// Tags:
// Backtracking

func shoppingOffers(price []int, special [][]int, needs []int) int {
	out := 0
	for i := range price {
		out += price[i] * needs[i]
	}
	dfs(special, price, needs, 0, 0, &out)
	return out
}

func dfs(special [][]int, price, needs []int, idx, cur int, out *int) {
	if idx == len(special) {
		return
	}
	dfs(special, price, needs, idx+1, cur, out)
	for i := range needs {
		if special[idx][i] > needs[i] {
			return
		}
	}
	for i := range needs {
		needs[i] -= special[idx][i]
	}
	cur += special[idx][len(price)]
	dfs(special, price, needs, idx, cur, out)
	dfs(special, price, needs, idx+1, cur, out)
	for i := range needs {
		cur += price[i] * needs[i]
	}
	if cur < *out {
		*out = cur
	}
	for i := range needs {
		needs[i] += special[idx][i]
	}
}
