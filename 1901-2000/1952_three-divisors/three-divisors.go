package main

func isThree(n int) bool {
	var cnt int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			cnt++
		}
		if cnt > 3 {
			return false
		}
	}
	return cnt == 3
}
