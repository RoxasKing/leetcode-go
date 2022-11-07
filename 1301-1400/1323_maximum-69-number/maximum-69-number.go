package main

// Difficulty:
// Easy

// Tags:
// Greedy

func maximum69Number(num int) int {
	a := []int{}
	for ; num > 0; num /= 10 {
		a = append(a, num%10)
	}
	o, ok := 0, false
	for i := len(a) - 1; i >= 0; i-- {
		if !ok && a[i] == 6 {
			ok, a[i] = true, 9
		}
		o = o*10 + a[i]
	}
	return o
}
