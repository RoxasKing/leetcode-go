package main

// Tags:
// Bit Manipulation

func reverseBits(num int) int {
	arr := []int{}
	cnt := 0
	x := uint(num)
	for i := 0; i <= 31; i++ {
		if (x>>i)&1 == 1 {
			cnt++
		} else {
			arr = append(arr, cnt, 0)
			cnt = 0
		}
	}
	arr = append(arr, cnt)
	out := 0
	for i := range arr {
		if arr[i] != 0 {
			out = Max(out, arr[i])
			continue
		}
		cnt := 1
		if i > 0 {
			cnt += arr[i-1]
		}
		if i < len(arr)-1 {
			cnt += arr[i+1]
		}
		out = Max(out, cnt)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
