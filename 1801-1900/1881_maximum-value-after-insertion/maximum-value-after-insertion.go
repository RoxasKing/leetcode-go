package main

// Tags:
// Greedy

func maxValue(n string, x int) string {
	flg := 1
	i := 0
	if n[0] == '-' {
		flg = -flg
		i++
	}
	for ; i < len(n); i++ {
		if flg*x > flg*int(n[i]-'0') {
			break
		}
	}
	return n[:i] + string('0'+byte(x)) + n[i:]
}
