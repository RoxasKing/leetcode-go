package main

// Tags:
// Greedy Algorithm
func candy(ratings []int) int {
	n := len(ratings)
	dp := make([]int, n)
	q := []int{}

	for i, rt := range ratings {
		if i-1 >= 0 && ratings[i-1] < rt {
			continue
		}
		if i+1 < n && ratings[i+1] < rt {
			continue
		}
		dp[i] = 1
		q = append(q, i)
	}

	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		if i-1 >= 0 && ratings[i-1] > ratings[i] && dp[i-1] < dp[i]+1 {
			dp[i-1] = dp[i] + 1
			q = append(q, i-1)
		}
		if i+1 < n && ratings[i+1] > ratings[i] && dp[i+1] < dp[i]+1 {
			dp[i+1] = dp[i] + 1
			q = append(q, i+1)
		}
	}

	out := 0
	for _, num := range dp {
		out += num
	}
	return out
}
