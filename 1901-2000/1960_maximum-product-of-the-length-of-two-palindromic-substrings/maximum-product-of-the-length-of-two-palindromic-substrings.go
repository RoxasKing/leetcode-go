package main

// Tags:
// Manacher

func maxProduct(s string) int64 {
	n := len(s)

	armLen := make([]int, n)
	f1 := make([]int, n)
	f2 := make([]int, n)
	for i := range f1 {
		f1[i] = 1
		f2[i] = 1
	}

	for i, j, right := 0, -1, -1; i < n; i++ {
		l, r := i, i
		if right >= i {
			iSym := j<<1 - i
			minArmLen := Min(armLen[iSym], right-i)
			l, r = i-minArmLen, i+minArmLen
		}
		for l-1 >= 0 && r+1 < n && s[l-1] == s[r+1] {
			l--
			r++
			f1[r] = Max(f1[r], r+1-l)
		}
		curArmLen := r - i
		armLen[i] = curArmLen
		if i+curArmLen > right {
			j, right = i, i+curArmLen
		}
	}

	for i, j, left := n-1, -1, n; i >= 0; i-- {
		l, r := i, i
		if left <= i {
			iSym := j<<1 - i
			minArmLen := Min(armLen[iSym], i-left)
			l, r = i-minArmLen, i+minArmLen
		}
		for l-1 >= 0 && r+1 < n && s[l-1] == s[r+1] {
			l--
			r++
			f2[l] = Max(f2[l], r+1-l)
		}
		curArmLen := r - i
		armLen[i] = curArmLen
		if i-curArmLen < left {
			j, left = i, i-curArmLen
		}
	}

	for i := 1; i < n; i++ {
		f1[i] = Max(f1[i], f1[i-1])
		f2[n-1-i] = Max(f2[n-1-i], f2[n-i])
	}

	var out int64
	for i := 0; i < n-1; i++ {
		out = Max64(out, int64(f1[i])*int64(f2[i+1]))
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
