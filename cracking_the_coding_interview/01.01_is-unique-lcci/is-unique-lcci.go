package main

// Tags:
// Bit Manipulation

func isUnique(astr string) bool {
	cur := 0
	for i := 0; i < len(astr); i++ {
		if cur&(1<<int(astr[i]-'a')) > 0 {
			return false
		}
		cur |= 1 << int(astr[i]-'a')
	}
	return true
}
