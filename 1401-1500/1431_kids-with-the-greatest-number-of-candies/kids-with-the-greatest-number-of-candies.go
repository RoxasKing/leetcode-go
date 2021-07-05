package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for _, candy := range candies {
		max = Max(max, candy)
	}
	n := len(candies)
	out := make([]bool, n)
	for i := 0; i < n; i++ {
		out[i] = candies[i]+extraCandies >= max
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
