package main

// Tags:
// Dynamic Programming
func splitArray(nums []int) int {
	n := len(nums)
	for i := 2; i < 1e6+1; i++ {
		dp[i] = n
		if minPrime[i] != 0 {
			continue
		}
		for j := i; j < 1e6+1; j += i {
			if minPrime[j] == 0 {
				minPrime[j] = i
			}
		}
	}

	out := 0
	for _, num := range nums {
		min := n
		for num > 1 {
			factor := minPrime[num]
			dp[factor] = Min(dp[factor], out+1)
			min = Min(dp[factor], min)

			num /= factor
		}
		out = min
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var minPrime [1e6 + 1]int
var dp [1e6 + 1]int
