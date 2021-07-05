package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	n := len(s)
	dict1 := [128]byte{}
	dict2 := [128]byte{}
	for i := 0; i < n; i++ {
		if dict1[s[i]] == 0 {
			dict1[s[i]] = t[i]
		} else if dict1[s[i]] != t[i] {
			return false
		}
		if dict2[t[i]] == 0 {
			dict2[t[i]] = s[i]
		} else if dict2[t[i]] != s[i] {
			return false
		}
	}
	return true
}
