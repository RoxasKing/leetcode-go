package main

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return toInt(firstWord)+toInt(secondWord) == toInt(targetWord)
}

func toInt(s string) int {
	out := 0
	for i := range s {
		out = out*10 + int(s[i]-'a')
	}
	return out
}
