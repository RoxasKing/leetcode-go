package main

// Tags:
// Hash + Congruence Theorem
func subarraysDivByK(A []int, K int) int {
	out := 0
	dict := make(map[int]int)
	dict[0] = 1
	sum := 0
	for _, a := range A {
		sum += a
		sum %= K
		if sum < 0 {
			sum += K
		}
		out += dict[sum]
		dict[sum]++
	}
	return out
}
