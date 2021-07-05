package main

func checkIfPangram(sentence string) bool {
	mark := [26]bool{}
	cnt := 0
	for i := range sentence {
		idx := int(sentence[i] - 'a')
		if mark[idx] {
			continue
		}
		cnt++
		mark[idx] = true
	}
	return cnt == 26
}
