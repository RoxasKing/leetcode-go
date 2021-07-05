package main

// Tags:
// Tags:
// Binary Search
// Hash

func longestCommonSubpath(n int, paths [][]int) int {
	l, r := 0, 1<<31-1
	for i := range paths {
		r = Min(r, len(paths[i]))
	}
	for l < r {
		m := l + (r-l+1)>>1
		if isValid(paths, m) {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}

func isValid(paths [][]int, size int) bool {
	ss := 1
	for i := 0; i < size; i++ {
		ss *= seed
	}
	freq := map[int]int{}
	for _, p := range paths {
		sum := 0
		added := map[int]bool{}
		for i := range p {
			sum *= seed
			sum += p[i]
			if i > size-1 {
				sum -= p[i-size] * ss
			}
			if i >= size-1 {
				if !added[sum] {
					freq[sum]++
					if freq[sum] == len(paths) {
						return true
					}
				}
				added[sum] = true
			}
		}
	}
	return false
}

var seed = int(1e6 + 7)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
