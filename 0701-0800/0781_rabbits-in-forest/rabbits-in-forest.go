package main

// Tags:
// Math + Hash
func numRabbits(answers []int) int {
	dict := make(map[int]int)
	for _, a := range answers {
		dict[a]++
	}
	out := 0
	for k, v := range dict {
		if k+1 >= v {
			out += k + 1
		} else {
			t := v / (k + 1)
			out += t * (k + 1)
			if v%(k+1) > 0 {
				out += k + 1
			}
		}
	}
	return out
}
