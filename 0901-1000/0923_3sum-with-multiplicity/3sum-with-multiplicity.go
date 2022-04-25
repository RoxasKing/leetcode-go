package main

// Difficulty:
// Medium

// Tags:
// Math
// Counting

func threeSumMulti(arr []int, target int) int {
	f := [101]int{}
	for _, a := range arr {
		f[a]++
	}
	o := 0
	for a := 0; a*3 < target; a++ {
		for b := Max(a+1, target-100-a); a+b*2 < target; b++ {
			o = (o + f[a]*f[b]*f[target-a-b]) % mod
		}
		if b := target - a*2; b < 101 && f[a] > 1 {
			o = (o + f[a]*(f[a]-1)/2*f[b]) % mod
		}
		if b := (target - a) / 2; (target-a)&1 == 0 && b < 101 && f[b] > 1 {
			o = (o + f[a]*f[b]*(f[b]-1)/2) % mod
		}
	}
	if a := target / 3; target%3 == 0 && f[a] > 2 {
		o = (o + f[a]*(f[a]-1)*(f[a]-2)/6) % mod
	}
	return o
}

const mod int = 1e9 + 7

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
