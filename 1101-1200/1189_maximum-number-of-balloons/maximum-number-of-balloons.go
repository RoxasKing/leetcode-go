package main

// Difficulty:
// Easy

// Tags:
// Hash

func maxNumberOfBalloons(text string) int {
	freq := [26]int{}
	for i := range text {
		freq[text[i]-'a']++
	}
	out := 1<<31 - 1
	out = Min(out, freq[0])
	out = Min(out, freq['b'-'a'])
	out = Min(out, freq['l'-'a']>>1)
	out = Min(out, freq['n'-'a'])
	out = Min(out, freq['o'-'a']>>1)
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
