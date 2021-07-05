package main

// Tags:
// Math

func threeEqualParts(arr []int) []int {
	cnt := 0
	for _, num := range arr {
		if num == 1 {
			cnt++
		}
	}
	if cnt%3 != 0 {
		return []int{-1, -1}
	}

	target := cnt / 3

	cnt, a, b := 0, -1, -1
	for i, num := range arr {
		if num == 1 {
			cnt++
		}

		if cnt < target {
			continue
		}
		cnt = 0

		if a == -1 {
			a = i
			continue
		}

		b = i

		num3 := 0
		for j := b + 1; j < len(arr); j++ {
			num3 = num3<<1 + arr[j]
			num3 %= 1e9 + 7
		}

		num1 := 0
		for j := 0; j <= a; j++ {
			num1 = num1<<1 + arr[j]
			num1 %= 1e9 + 7
		}
		for arr[a+1] == 0 && num1 != num3 {
			a++
			num1 <<= 1
			num1 %= 1e9 + 7
		}

		if num1 != num3 {
			return []int{-1, -1}
		}

		num2 := 0
		for j := a + 1; j <= b; j++ {
			num2 = num2<<1 + arr[j]
			num2 %= 1e9 + 7
		}
		for arr[b+1] == 0 && num2 != num3 {
			b++
			num2 <<= 1
			num2 %= 1e9 + 7
		}

		if num2 != num3 {
			return []int{-1, -1}
		}

		b++
		break
	}

	return []int{a, b}
}
