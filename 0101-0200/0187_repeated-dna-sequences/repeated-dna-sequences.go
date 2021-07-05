package main

// Tags:
// Rabin-Karp
func findRepeatedDnaSequences(s string) []string {
	n, L := len(s), 10
	if n <= L {
		return nil
	}
	dict := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	nums := make([]int, n)
	for i := range nums {
		nums[i] = dict[s[i]]
	}
	base := 4
	baseL := Pow(base, L)
	var hash int
	for i := 0; i < L; i++ {
		hash = hash*base + nums[i]
	}
	mark := map[int]bool{hash: true}
	save := make(map[int]bool)
	var out []string
	for i := 1; i < n+1-L; i++ {
		hash = hash*base - nums[i-1]*baseL + nums[i-1+L]
		if mark[hash] && !save[hash] {
			out = append(out, s[i:i+L])
			save[hash] = true
		}
		mark[hash] = true
	}
	return out
}

func Pow(base, t int) int {
	out := 1
	for t > 0 {
		if t&1 == 1 {
			out *= base
		}
		base *= base
		t >>= 1
	}
	return out
}
