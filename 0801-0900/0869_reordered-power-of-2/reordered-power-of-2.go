package main

// Difficulty:
// Medium

// Tags:
// Counting

func reorderedPowerOf2(n int) bool {
	f0 := [10]int{}
	for ; n > 0; n /= 10 {
		f0[n%10]++
	}
	for i := 1; i <= 1e9; i <<= 1 {
		f1 := [10]int{}
		for x := i; x > 0; x /= 10 {
			f1[x%10]++
		}
		ok := true
		for i := 0; i < 10 && ok; i++ {
			ok = f0[i] == f1[i]
		}
		if ok {
			return true
		}
	}
	return false
}
