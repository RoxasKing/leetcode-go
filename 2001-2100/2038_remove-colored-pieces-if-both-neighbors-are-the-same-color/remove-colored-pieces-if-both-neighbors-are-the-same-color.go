package main

// Difficulty:
// Medium

// Tags:
// Game Theory

func winnerOfGame(colors string) bool {
	sum, flg, cnt := 0, 1, 0
	for i := range colors {
		if c := colors[i]; c == 'A' && flg == 1 || c == 'B' && flg == -1 {
			cnt++
			continue
		}
		sum += flg * Max(0, cnt-2)
		flg, cnt = -flg, 1
	}
	sum += flg * Max(0, cnt-2)
	return sum > 0
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
