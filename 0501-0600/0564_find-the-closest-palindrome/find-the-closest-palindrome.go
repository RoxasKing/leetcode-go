package main

// Difficulty:
// Hard

func nearestPalindromic(n string) string {
	arr := []byte(n)
	for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
		if arr[l] != arr[r] {
			arr[r] = arr[l]
		}
	}
	out := ""
	compareAndSwap := func(tmp string) {
		if len(out) == 0 {
			out = tmp
			return
		}
		d1, d2 := dist(n, out), dist(n, tmp)
		if len(d1) > len(d2) ||
			len(d1) == len(d2) && d1 > d2 ||
			len(d1) == len(d2) && d1 == d2 && (len(out) > len(tmp) || len(out) == len(tmp) && out > tmp) {
			out = tmp
		}
	}
	for x := 0; x <= 9; x++ {
		arr[len(arr)>>1] = byte(x) + '0'
		if len(arr)&1 == 0 {
			arr[len(arr)>>1-1] = byte(x) + '0'
		}
		if string(arr) != n {
			compareAndSwap(string(arr))
		}
	}
	if len(out) == 1 {
		return out
	}
	arr = make([]byte, len(n)-1)
	for i := range arr {
		arr[i] = '9'
	}
	if tmp := string(arr); tmp != n {
		compareAndSwap(tmp)
	}
	arr = make([]byte, len(n)+1)
	arr[0], arr[len(arr)-1] = '1', '1'
	for i := 1; i < len(arr)-1; i++ {
		arr[i] = '0'
	}
	if tmp := string(arr); tmp != n {
		compareAndSwap(tmp)
	}
	return out
}

func dist(a, b string) string {
	if len(a) < len(b) || len(a) == len(b) && a < b {
		a, b = b, a
	}
	arr := []byte{}
	for i, j, del := len(a)-1, len(b)-1, 0; i >= 0; i-- {
		x := int(a[i] - '0')
		y := del
		if j >= 0 {
			y += int(b[j] - '0')
			j--
		}
		if x < y {
			arr = append(arr, byte(x+10-y)+'0')
			del = 1
		} else {
			arr = append(arr, byte(x-y)+'0')
			del = 0
		}
	}
	for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
		arr[l], arr[r] = arr[r], arr[l]
	}
	for ; len(arr) > 1 && arr[0] == '0'; arr = arr[1:] {
	}
	return string(arr)
}
