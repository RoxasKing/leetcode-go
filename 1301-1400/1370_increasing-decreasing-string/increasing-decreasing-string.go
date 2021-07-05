package main

func sortString(s string) string {
	count := [26]int{}
	for i := range s {
		count[s[i]-'a']++
	}
	n := len(s)
	out := make([]byte, 0, n)
	for len(out) < n {
		for i := 0; i < 26; i++ {
			if count[i] > 0 {
				out = append(out, byte(i)+'a')
				count[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if count[i] > 0 {
				out = append(out, byte(i)+'a')
				count[i]--
			}
		}
	}
	return string(out)
}
