package main

func game(guess []int, answer []int) int {
	out := 0
	for i := range guess {
		if guess[i] == answer[i] {
			out++
		}
	}
	return out
}
