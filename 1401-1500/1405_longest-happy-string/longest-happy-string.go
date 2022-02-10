package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Greedy

func longestDiverseString(a int, b int, c int) string {
	chs := make([]byte, 0, a+b+c)
	arr := []struct {
		freq   int
		letter byte
	}{{a, 'a'}, {b, 'b'}, {c, 'c'}}
	for finished := false; !finished; {
		sort.Slice(arr, func(i, j int) bool { return arr[i].freq > arr[j].freq })
		finished = true
		for i, e := range arr {
			if e.freq == 0 {
				break
			}
			if m := len(chs); m > 1 && chs[m-2] == e.letter && chs[m-1] == e.letter {
				continue
			}
			finished = false
			chs = append(chs, e.letter)
			arr[i].freq--
			break
		}
	}
	return string(chs)
}
