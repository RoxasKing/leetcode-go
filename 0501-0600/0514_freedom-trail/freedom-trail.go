package main

// Tags:
// Dynamic Programming
func findRotateSteps(ring string, key string) int {
	m, n := len(ring), len(key)
	dict := [26][]int{}
	for i := 0; i < m; i++ {
		idx := ring[i] - 'a'
		dict[idx] = append(dict[idx], i)
	}
	solutions := make([]solution, len(dict[key[0]-'a']))
	for i, idx := range dict[key[0]-'a'] {
		solutions[i] = solution{index: idx, steps: 1 + Min(idx, m-idx)}
	}
	for i := 1; i < n; i++ {
		newPairs := make([]solution, len(dict[key[i]-'a']))
		for j, next := range dict[key[i]-'a'] {
			minStep := 1<<31 - 1
			for _, p := range solutions {
				nextStep := Min(Abs(next-p.index), m-Abs(next-p.index)) + 1
				minStep = Min(minStep, p.steps+nextStep)
			}
			newPairs[j] = solution{index: next, steps: minStep}
		}
		solutions = newPairs
	}
	out := 1<<31 - 1
	for _, p := range solutions {
		out = Min(out, p.steps)
	}
	return out
}

type solution struct {
	index int
	steps int
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
