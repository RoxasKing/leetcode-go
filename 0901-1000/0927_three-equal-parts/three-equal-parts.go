package main

// Difficulty:
// Hard

// Tags:
// Math
// Two Pointers

func threeEqualParts(arr []int) []int {
	const mod int = 1e9 + 7
	cnt := 0
	for _, x := range arr {
		if x == 1 {
			cnt++
		}
	}
	if cnt%3 != 0 {
		return []int{-1, -1}
	}
	if cnt == 0 {
		return []int{0, len(arr) - 1}
	}
	c, i, j := 0, -1, -1
	for k, x := range arr {
		if x == 1 {
			if c++; c < cnt/3 {
				continue
			}
		}
		c = 0
		if i == -1 {
			i = k
		} else {
			j = k
			break
		}
	}
	num3 := 0
	for _, x := range arr[j+1:] {
		num3 = (num3<<1 + x) % mod
	}
	num1 := 0
	for _, x := range arr[:i+1] {
		num1 = (num1<<1 + x) % mod
	}
	for ; arr[i+1] == 0 && num1 != num3; i++ {
		num1 = (num1 << 1) % mod
	}
	if num1 != num3 {
		return []int{-1, -1}
	}
	num2 := 0
	for _, x := range arr[i+1 : j+1] {
		num2 = (num2<<1 + x) % mod
	}
	for ; arr[j+1] == 0 && num2 != num3; j++ {
		num2 = (num2 << 1) % mod
	}
	if num2 != num3 {
		return []int{-1, -1}
	}
	return []int{i, j + 1}
}
