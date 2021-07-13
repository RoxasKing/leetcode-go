package main

// Tags:
// Greedy Algorithm

func sumGame(num string) bool {
	var cnt, sum int
	n := len(num)
	for i := 0; i < n>>1; i++ {
		if num[i] == '?' {
			cnt++
		} else {
			sum += int(num[i] - '0')
		}
		if num[n-1-i] == '?' {
			cnt--
		} else {
			sum -= int(num[n-1-i] - '0')
		}
	}

	if cnt < 0 {
		cnt = -cnt
		sum = -sum
	}

	if cnt&1 == 0 {
		return sum+cnt/2*9 != 0
	}
	return sum+cnt/2*9 != 0 || sum+9+cnt/2*9 != 0
}
