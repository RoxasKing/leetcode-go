package main

// Tags:
// Stack
// Dynamic Programming

func minOperationsToFlip(expression string) int {
	stk := IntStack{} // record previous '(', arr's lenght
	arr := [][2]int{} // [minOpTo 0, minOpTo 1]
	ops := []byte{}
	for i := range expression {
		op := expression[i]
		switch op {
		case '(':
			stk.Push(len(arr))
		case ')':
			start := stk.Pop()
			out := minOps(arr[start:], ops[start:])
			arr = arr[:start]
			ops = ops[:start]
			arr = append(arr, out)
		case '&', '|':
			ops = append(ops, op)
		case '0':
			arr = append(arr, [2]int{0, 1})
		case '1':
			arr = append(arr, [2]int{1, 0})
		}
	}

	out := minOps(arr, ops)

	if out[0] == 0 {
		return out[1]
	}
	return out[0]
}

func minOps(arr [][2]int, ops []byte) [2]int {
	n := len(arr)
	f := make([][2]int, n) // the min num of ops required to 0 or 1
	f[0] = arr[0]
	for i := 1; i < n; i++ {
		to0, to1 := arr[i][0], arr[i][1]
		if ops[i-1] == '&' {
			f[i][0] = Min(f[i-1][0]+to0, Min(f[i-1][1]+to0, f[i-1][0]+to1))   // 0 & 0 or 1 & 0 or 0 & 1
			f[i][1] = Min(f[i-1][1]+to1, Min(f[i-1][1]+to0, f[i-1][0]+to1)+1) // 1 & 1 or 1 | 0 or 0 | 1
		} else {
			f[i][0] = Min(f[i-1][0]+to0, Min(f[i-1][1]+to0, f[i-1][0]+to1)+1) // 0 | 0 or 0 & 1 or 1 & 0
			f[i][1] = Min(f[i-1][1]+to1, Min(f[i-1][1]+to0, f[i-1][0]+to1))   // 1 | 1 or 1 | 0 or 0 | 1
		}
	}
	return f[n-1]
}

type IntStack []int

func (s IntStack) Len() int    { return len(s) }
func (s IntStack) Top() int    { return s[s.Len()-1] }
func (s *IntStack) Push(x int) { *s = append(*s, x) }
func (s *IntStack) Pop() int {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
