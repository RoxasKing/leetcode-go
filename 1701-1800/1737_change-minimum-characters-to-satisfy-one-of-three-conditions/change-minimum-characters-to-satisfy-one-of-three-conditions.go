package main

// Tags:
// Greedy

func minCharacters(a string, b string) int {
	freqA := [26]int{}
	freqB := [26]int{}

	for i := range a {
		freqA[a[i]-'a']++
	}
	for i := range b {
		freqB[b[i]-'a']++
	}

	out := len(a) + len(b)
	cntA, cntB := 0, 0
	for i := 0; i < 26; i++ {
		cntA += freqA[i]
		cntB += freqB[i]
		if i < 25 {
			out = Min(out, cntB+len(a)-cntA)
			out = Min(out, cntA+len(b)-cntB)
		}
		out = Min(out, len(a)+len(b)-(freqA[i]+freqB[i]))
	}

	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
