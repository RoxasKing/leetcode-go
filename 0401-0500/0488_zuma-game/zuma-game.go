package main

// Difficulty:
// Hard

// Tags:
// BFS

func findMinStep(board string, hand string) int {
	mp := map[rune]int8{'R': 0, 'Y': 1, 'B': 2, 'G': 3, 'W': 4}
	a := []pair{}
	cur, cnt := int8(0), int8(0)
	for _, c := range board {
		if cur == mp[c] {
			cnt++
			continue
		}
		if cnt > 0 {
			a = append(a, pair{cur, cnt})
		}
		cur, cnt = mp[c], 1
	}
	a = append(a, pair{cur, cnt})
	h := [5]int8{}
	for _, c := range hand {
		h[mp[c]]++
	}
	for q := []state{{a, h, 0}}; len(q) > 0; q = q[1:] {
		e := q[0]
		if e.h[0] == 0 && e.h[1] == 0 && e.h[2] == 0 && e.h[3] == 0 && e.h[4] == 0 {
			continue
		}
		c := e.c + 1
		for i, p := range e.a {
			h := e.h
			if h[p.k] > 0 {
				a := make([]pair, len(e.a))
				copy(a, e.a)
				if a[i].v++; a[i].v >= 3 {
					a = a[:i]
					for _, t := range e.a[i+1:] {
						if len(a) == 0 || a[len(a)-1].k != t.k {
							a = append(a, t)
							continue
						}
						if a[len(a)-1].v += t.v; a[len(a)-1].v >= 3 {
							a = a[:len(a)-1]
						}
					}
					if len(a) == 0 {
						return int(c)
					}
				}
				h[p.k]-- // Attention!! Don't use h[P.K]++ after append, or the running speed will be very slow.
				q = append(q, state{a, h, c})
			}
			if p.v != 2 {
				continue
			}
			for k := int8(0); k < 5; k++ {
				if h[k] == 0 || p.k == k {
					continue
				}
				a := make([]pair, 0, len(e.a)+2)
				a = append(a, e.a[:i]...)
				a = append(a, pair{p.k, 1}, pair{k, 1}, pair{p.k, 1})
				a = append(a, e.a[i+1:]...)
				h[k]--
				q = append(q, state{a, h, c})
				h[k]++
			}
		}
	}
	return -1
}

type pair struct{ k, v int8 }
type state struct {
	a []pair
	h [5]int8
	c int8
}
