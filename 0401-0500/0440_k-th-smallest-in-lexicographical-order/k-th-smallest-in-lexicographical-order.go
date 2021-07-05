package main

func findKthNumber(n int, k int) int {
	out := 1
	k--
	for k > 0 {
		step, head, tail := 0, out, out+1
		for head <= n {
			step += Min(n+1, tail) - head
			head *= 10
			tail *= 10
		}
		if step > k {
			out *= 10
			k--
		} else {
			out++
			k -= step
		}
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
