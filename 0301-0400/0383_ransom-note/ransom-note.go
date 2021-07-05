package main

func canConstruct(ransomNote string, magazine string) bool {
	cnt := [26]int{}
	for i := range magazine {
		cnt[magazine[i]-'a']++
	}
	for i := range ransomNote {
		idx := ransomNote[i] - 'a'
		cnt[idx]--
		if cnt[idx] < 0 {
			return false
		}
	}
	return true
}
