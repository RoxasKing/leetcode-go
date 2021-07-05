package main

func oneEditAway(first string, second string) bool {
	n1, n2 := len(first), len(second)
	if n1-n2 > 1 || n1-n2 < -1 {
		return false
	}
	for i := 0; i < n1 && i < n2; i++ {
		if first[i] != second[i] {
			return first[i+1:] == second[i+1:] || first[i+1:] == second[i:] || first[i:] == second[i+1:]
		}
	}
	return true
}
