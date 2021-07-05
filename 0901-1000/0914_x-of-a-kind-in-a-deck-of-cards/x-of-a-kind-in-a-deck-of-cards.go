package main

// Tags:
// Math
func hasGroupsSizeX(deck []int) bool {
	cnt := make(map[int]int)
	for _, num := range deck {
		cnt[num]++
	}
	prev := -1
	for _, v := range cnt {
		if prev != -1 {
			prev = Gcd(prev, v)
		} else {
			prev = v
		}
	}
	return prev >= 2
}

func Gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
