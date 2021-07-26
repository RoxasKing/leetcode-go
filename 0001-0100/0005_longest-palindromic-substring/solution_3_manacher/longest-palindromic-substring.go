package main

// Tags:
// Manacher

func longestPalindrome(s string) string {
	t := "#"
	for i := range s {
		t += s[i:i+1] + "#"
	}
	s = t

	n := len(s)
	armLen := make([]int, n)
	head, tail := 0, -1

	for i, j, right := 0, -1, -1; i < n; i++ {
		curArmLen := 0
		if right >= i {
			iSym := j<<1 - i
			minArmLen := Min(armLen[iSym], right-i)
			curArmLen = expand(s, i-minArmLen, i+minArmLen)
		} else {
			curArmLen = expand(s, i, i)
		}
		armLen[i] = curArmLen
		if i+curArmLen > right {
			j = i
			right = i + curArmLen
		}
		if curArmLen<<1+1 > tail-head {
			head, tail = i-curArmLen, i+curArmLen
		}
	}

	arr := make([]byte, 0, (tail-head)>>1)
	for i := head; i <= tail; i++ {
		if s[i] != '#' {
			arr = append(arr, s[i])
		}
	}
	return string(arr)
}

func expand(s string, l, r int) int {
	for l-1 >= 0 && r+1 < len(s) && s[l-1] == s[r+1] {
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
