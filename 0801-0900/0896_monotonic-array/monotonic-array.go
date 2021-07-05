package main

func isMonotonic(A []int) bool {
	inc := []int{}
	dec := []int{}
	for _, a := range A {
		if len(inc) <= 1 {
			if len(inc) == 1 && inc[0] <= a {
				inc = inc[:0]
			}
			inc = append(inc, a)
		}
		if len(dec) <= 1 {
			if len(dec) == 1 && dec[0] >= a {
				dec = dec[:0]
			}
			dec = append(dec, a)
		}
		if len(inc) > 1 && len(dec) > 1 {
			return false
		}
	}
	return true
}
