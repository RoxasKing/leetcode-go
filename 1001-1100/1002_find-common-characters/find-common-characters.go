package main

func commonChars(A []string) []string {
	dict := [26]int{}
	for i := range A[0] {
		dict[A[0][i]-'a']++
	}
	for i := 1; i < len(A); i++ {
		newDict := [26]int{}
		for j := range A[i] {
			index := A[i][j] - 'a'
			if dict[index] > 0 {
				newDict[index]++
				dict[index]--
			}
		}
		dict = newDict
	}
	out := []string{}
	for i := range dict {
		for dict[i] > 0 {
			out = append(out, string(byte(i+'a')))
			dict[i]--
		}
	}
	return out
}
