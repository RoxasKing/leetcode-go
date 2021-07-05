package main

// Tags:
// Dynamic Programming
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if n1 == 0 {
		return 0
	}
	s1cnt, index, s2cnt := 0, 0, 0
	recall := make(map[int][]int)
	var preLoop, inLoop []int
	for {
		s1cnt++
		for i := range s1 {
			if s1[i] == s2[index] {
				index++
				if index == len(s2) {
					s2cnt, index = s2cnt+1, 0
				}
			}
		}
		if s1cnt == n1 {
			return s2cnt / n2
		}
		if _, ok := recall[index]; ok {
			preLoop = recall[index]
			inLoop = []int{s1cnt - recall[index][0], s2cnt - recall[index][1]}
			break
		} else {
			recall[index] = []int{s1cnt, s2cnt}
		}
	}
	ans := preLoop[1] + (n1-preLoop[0])/inLoop[0]*inLoop[1]
	rest := (n1 - preLoop[0]) % inLoop[0]
	for i := 0; i < rest; i++ {
		for j := range s1 {
			if s1[j] == s2[index] {
				index++
				if index == len(s2) {
					ans, index = ans+1, 0
				}
			}
		}
	}
	return ans / n2
}
