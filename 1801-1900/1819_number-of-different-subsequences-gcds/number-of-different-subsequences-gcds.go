package main

// Tags:
// Math
func countDifferentSubsequenceGCDs(nums []int) int {
	out := 0
	dict := make(map[int]bool)
	for _, num := range nums {
		if !dict[num] {
			out++
		}
		dict[num] = true
	}
	for i := 1; i <= 1e5; i++ {
		if dict[i] {
			continue
		}
		gcd := -1
		for j := i << 1; j <= 2*1e5; j += i {
			if !dict[j] {
				continue
			}
			if gcd != -1 {
				gcd = Gcd(gcd, j)
			} else {
				gcd = j
			}
			if gcd == i {
				out++
				break
			}
		}
	}
	return out
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
