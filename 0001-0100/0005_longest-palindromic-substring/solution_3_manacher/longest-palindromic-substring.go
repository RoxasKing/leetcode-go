package main

// Tags:
// Manacher

func longestPalindrome(s string) string {
	t := append(make([]byte, 0, len(s)<<1+1), '#')
	for i := range s {
		t = append(t, s[i], '#')
	}

	n := len(t)
	armLen := make([]int, n)
	head, tail := 0, -1

	for i, j, right := 0, -1, -1; i < n; i++ {
		curArmLen := 0
		if i > right {
			curArmLen = expand(t, i, i)
		} else {
			iSym := j<<1 - i
			minArmLen := Min(armLen[iSym], right-i)
			curArmLen = expand(t, i-minArmLen, i+minArmLen)
		}
		armLen[i] = curArmLen
		if i+curArmLen > right {
			j = i
			right = i + curArmLen
		}
		if curArmLen<<1+1 > tail+1-head {
			head, tail = i-curArmLen, i+curArmLen
		}
	}

	chs := make([]byte, 0, (tail-head)>>1)
	for i := head; i <= tail; i++ {
		if t[i] != '#' {
			chs = append(chs, t[i])
		}
	}
	return string(chs)
}

func expand(x []byte, l, r int) int {
	for 0 <= l-1 && r+1 < len(x) && x[l-1] == x[r+1] {
		l--
		r++
	}
	return (r - l) >> 1
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
