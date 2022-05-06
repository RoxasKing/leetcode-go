package main

// Difficulty:
// Medium

// Tags:
// Stack

func removeDuplicates(s string, k int) string {
	stk := []pair{}
	top := func() int { return len(stk) - 1 }
	cur, cnt := s[0], 1
	handle := func() {
		if cnt %= k; cnt == 0 {
			return
		}
		if len(stk) > 0 && stk[top()].k == cur {
			stk[top()].v += cnt
			if stk[top()].v %= k; stk[top()].v == 0 {
				stk = stk[:top()]
			}
		} else {
			stk = append(stk, pair{cur, cnt})
		}
	}
	for i := 1; i < len(s); i++ {
		if cur == s[i] {
			cnt++
			continue
		}
		handle()
		cur, cnt = s[i], 1
	}
	handle()
	arr := []byte{}
	for _, p := range stk {
		c := p.k
		for i := 0; i < p.v; i++ {
			arr = append(arr, c)
		}
	}
	return string(arr)
}

type pair struct {
	k byte
	v int
}
