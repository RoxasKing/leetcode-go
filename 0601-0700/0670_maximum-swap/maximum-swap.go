package main

// Tags:
// Greedy
// Math

func maximumSwap(num int) int {
	a := []int{}
	for ; num > 0; num /= 10 {
		a = append(a, num%10)
	}
	for i := len(a) - 1; i > 0; i-- {
		j := 0
		for k := 1; k < i; k++ {
			if a[j] < a[k] {
				j = k
			}
		}
		if a[i] < a[j] {
			a[i], a[j] = a[j], a[i]
			break
		}
	}
	o := 0
	for i := len(a) - 1; i >= 0; i-- {
		o = o*10 + a[i]
	}
	return o
}
