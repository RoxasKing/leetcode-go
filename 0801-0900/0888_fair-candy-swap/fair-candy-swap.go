package main

func fairCandySwap(A []int, B []int) []int {
	sumA, sumB := 0, 0
	existed := map[int]bool{}
	for _, num := range A {
		sumA += num
	}
	for _, num := range B {
		sumB += num
		existed[num] = true
	}
	diff := (sumB - sumA) >> 1
	for _, num := range A {
		if existed[num+diff] {
			return []int{num, num + diff}
		}
	}
	return nil
}
