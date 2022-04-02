package main

// Difficulty:
// Hard

// Tags:
// Greedy

func strongPasswordChecker(password string) int {
	l, u, d := 0, 0, 0
	for i := range password {
		x := password[i]
		if 'a' <= x && x <= 'z' {
			l = 1
		} else if 'A' <= x && x <= 'Z' {
			u = 1
		} else if '0' <= x && x <= '9' {
			d = 1
		}
	}
	k, n := l+u+d, len(password)
	if n < 6 {
		return Max(3-k, 6-n)
	}
	if n < 21 {
		sum, cnt, cur := 0, 0, '$'
		for _, c := range password {
			if c == cur {
				cnt++
				continue
			}
			sum += cnt / 3
			cnt, cur = 1, c
		}
		sum += cnt / 3
		return Max(3-k, sum)
	}
	chg, del, del2, cnt, cur := 0 /* chg */, n-20 /* del */, 0 /* del2 */, 0 /* cnt */, '$' /* cur */
	decDelTimes := func() {
		if del > 0 && cnt > 2 {
			if res := cnt % 3; res == 0 {
				chg--
				del--
			} else if res == 1 {
				del2++
			}
		}
		chg += cnt / 3
	}
	for _, c := range password {
		if c == cur {
			cnt++
			continue
		}
		decDelTimes()
		cnt, cur = 1, c
	}
	decDelTimes()
	d2 := Min(chg, Min(del2, del/2))
	chg -= d2
	del -= d2 * 2
	d3 := Min(chg, del/3)
	chg -= d3
	return Max(3-k, chg) + n - 20
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
