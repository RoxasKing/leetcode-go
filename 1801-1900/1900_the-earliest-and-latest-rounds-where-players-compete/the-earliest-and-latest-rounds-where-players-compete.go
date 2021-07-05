package main

// Tags:
// Dynamic Porgramming

func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
	if firstPlayer > secondPlayer {
		firstPlayer, secondPlayer = secondPlayer, firstPlayer
	}

	return dp(n, firstPlayer, secondPlayer)
}

func dp(n, f, s int) []int {
	if F[n][f][s] > 0 {
		return []int{F[n][f][s], G[n][f][s]}
	}

	if f+s == n+1 {
		return []int{1, 1}
	}

	if f+s > n+1 {
		out := dp(n, n+1-s, n+1-f)
		F[n][f][s], G[n][f][s] = out[0], out[1]
		return out
	}

	min, max := 1<<31-1, -1<<31
	halfN := (n + 1) >> 1

	if s <= halfN {
		for i := 0; i < f; i++ {
			for j := 0; j < s-f; j++ {
				out := dp(halfN, i+1, i+j+2)
				min = Min(min, out[0])
				max = Max(max, out[1])
			}
		}
	} else {
		sPrime := n + 1 - s
		mid := (n - 2*sPrime + 1) >> 1
		for i := 0; i < f; i++ {
			for j := 0; j < sPrime-f; j++ {
				out := dp(halfN, i+1, i+j+mid+2)
				min = Min(min, out[0])
				max = Max(max, out[1])
			}
		}
	}

	F[n][f][s] = min + 1
	G[n][f][s] = max + 1
	return []int{F[n][f][s], G[n][f][s]}
}

var F = [30][30][30]int{}
var G = [30][30][30]int{}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
