package main

// Difficulty:
// Medium

// Tags:
// Recursion

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := countAndSay(n - 1)
	a := []rune{}
	cur, cnt := rune(s[0]), 1
	for _, c := range s[1:] {
		if c == cur {
			cnt++
			continue
		}
		a = append(a, rune(cnt)+'0', cur)
		cur, cnt = c, 1
	}
	a = append(a, rune(cnt)+'0', cur)
	return string(a)
}
