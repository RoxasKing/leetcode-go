package main

// Difficulty:
// Medium

// Tags:
// Sliding Window

func maxConsecutiveAnswers(answerKey string, k int) int {
	out := 0
	for l, r, c := 0, 0, 0; r < len(answerKey); r++ {
		if answerKey[r] == 'T' {
			c++
		}
		for ; c > k && r+1-l-c > k; l++ {
			if answerKey[l] == 'T' {
				c--
			}
		}
		if out < r+1-l {
			out = r + 1 - l
		}
	}
	return out
}
