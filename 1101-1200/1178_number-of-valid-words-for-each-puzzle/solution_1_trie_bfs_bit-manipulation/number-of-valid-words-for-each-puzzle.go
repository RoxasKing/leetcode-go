package main

// Difficulty:
// Hard

// Tags:
// Trie
// BFS
// Bit Manipulation

func findNumOfValidWords(words []string, puzzles []string) []int {
	t := &Trie{}
	for _, word := range words {
		t.Insert(word)
	}
	out := make([]int, len(puzzles))
	for i := range puzzles {
		has := [26]bool{}
		ks := []int{}
		for j := range puzzles[i] {
			idx := int(puzzles[i][j] - 'a')
			if !has[idx] {
				ks = append(ks, idx)
				has[idx] = true
			}
		}
		for q := []*Trie{t}; len(q) > 0; q = q[1:] {
			t := q[0]
			if t.freq > 0 && t.code&(1<<ks[0]) != 0 {
				out[i] += t.freq
			}
			if t.mask&(1<<ks[0]) == 0 {
				continue
			}
			for _, k := range ks {
				if t.next[k] != nil {
					q = append(q, t.next[k])
				}
			}
		}
	}
	return out
}

type Trie struct {
	next [26]*Trie
	mask int
	freq int
	code int
}

func (t *Trie) Insert(word string) {
	mask := 0
	for i := range word {
		mask |= 1 << int(word[i]-'a')
	}
	for i := range word {
		t.mask |= mask
		idx := word[i] - 'a'
		if t.next[idx] == nil {
			t.next[idx] = &Trie{}
		}
		t = t.next[idx]
	}
	t.freq++
	t.code = mask
}
