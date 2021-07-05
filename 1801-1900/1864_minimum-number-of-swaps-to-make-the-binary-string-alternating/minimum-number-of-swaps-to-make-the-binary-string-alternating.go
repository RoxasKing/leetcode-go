package main

// Tags:
// Math + Greedy Algorithm

func minSwaps(s string) int {
	cnt := 0
	for _, ch := range s {
		if ch == '0' {
			cnt++
		} else {
			cnt--
		}
	}

	if Abs(cnt) > 1 {
		return -1
	}

	cur := 0
	if cnt < 0 {
		cur = 1
	}
	diff := 0
	for i := range s {
		if int(s[i]-'0') != cur {
			diff++
		}
		cur ^= 1
	}
	diff >>= 1

	if cnt != 0 {
		return diff
	}

	out := diff
	cur, diff = 1, 0
	for i := range s {
		if int(s[i]-'0') != cur {
			diff++
		}
		cur ^= 1
	}
	diff >>= 1

	out = Min(out, diff)

	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
