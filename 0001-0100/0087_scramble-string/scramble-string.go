package main

// Tags:
// Dynamic Programming + DFS
func isScramble(s1 string, s2 string) bool {
	valid := make(map[string]bool)
	return dp(valid, s1, s2)
}

func dp(valid map[string]bool, s1, s2 string) bool {
	if s1 > s2 {
		s1, s2 = s2, s1
	}
	key := s1 + s2
	if val, ok := valid[key]; ok {
		return val
	}
	if s1 == s2 {
		valid[key] = true
		return true
	}
	cnt := make([]int, 128)
	for i := range s1 {
		cnt[s1[i]]++
	}
	for i := range s2 {
		if cnt[s2[i]]--; cnt[s2[i]] < 0 {
			valid[key] = false
			return false
		}
	}
	n := len(s1)
	for i := 1; i < n; i++ {
		if dp(valid, s1[:i], s2[:i]) && dp(valid, s1[i:], s2[i:]) ||
			dp(valid, s1[:i], s2[n-i:]) && dp(valid, s1[i:], s2[:n-i]) {
			valid[key] = true
			return true
		}
	}
	valid[key] = false
	return false
}
