package main

// Difficulty:
// Medium

// Tags:
// Stack

func lengthLongestPath(input string) int {
	type path struct {
		lvl, cnt int
		isf      bool
	}
	stk := []path{}
	top := func() int { return len(stk) - 1 }
	o := 0
	n := len(input)
	for i := 0; i < n; {
		lvl := 0
		for ; i < n && input[i] == '\t'; i++ {
			lvl++
		}
		isf := false
		j := i
		for ; j < n && input[j] != '\n'; j++ {
			if input[j] == '.' {
				isf = true
			}
		}
		for len(stk) > 0 && stk[top()].lvl >= lvl {
			stk = stk[:top()]
		}
		cnt := j - i
		if len(stk) > 0 {
			cnt += stk[top()].cnt + 1
		}
		stk = append(stk, path{lvl, cnt, isf})
		if e := stk[top()]; e.isf && o < e.cnt {
			o = e.cnt
		}
		i = j + 1
	}
	return o
}
