package main

// Difficulty:
// Easy

// Tags:
// Hash

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	cnt := 0
	freq1 := [26]int{}
	freq2 := [26]int{}
	for i := range s {
		if s[i] != goal[i] {
			cnt++
			if cnt > 2 {
				return false
			}
		}
		freq1[s[i]-'a']++
		freq2[goal[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}
	if cnt == 2 {
		return true
	}
	for i := 0; i < 26; i++ {
		if freq1[i] > 1 {
			return true
		}
	}
	return false
}
