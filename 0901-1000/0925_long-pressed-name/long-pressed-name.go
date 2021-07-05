package main

// Tags:
// Two Pointers
func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		if name[i] != typed[j] {
			return false
		}
		ch := name[i]
		lCnt := 0
		for i < len(name) && name[i] == ch {
			lCnt++
			i++
		}
		rCnt := 0
		for j < len(typed) && typed[j] == ch {
			rCnt++
			j++
		}
		if lCnt > rCnt {
			return false
		}
	}
	return i == len(name) && j == len(typed)
}
