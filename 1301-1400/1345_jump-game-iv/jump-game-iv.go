package main

// Difficulty:
// Hard

// Tags:
// Hash
// BFS

func minJumps(arr []int) int {
	h := map[int][]int{}
	for i, num := range arr {
		h[num] = append(h[num], i)
	}
	n := len(arr)
	f0, f1 := make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		f0[i] = 1<<31 - 1
		f1[n-1-i] = 1<<31 - 1
	}
	for q0, q1 := []int{0}, []int{n - 1}; len(q0) > 0; q0, q1 = q0[1:], q1[1:] {
		a := q0[0]
		step0 := f0[a] + 1
		if a+1 < n && step0 < f0[a+1] {
			f0[a+1] = step0
			q0 = append(q0, a+1)
		}
		if a-1 >= 0 && step0 < f0[a-1] {
			f0[a-1] = step0
			q0 = append(q0, a-1)
		}
		for _, b := range h[arr[a]] {
			if step0 < f0[b] {
				f0[b] = step0
				q0 = append(q0, b)
			}
		}
		if f0[n-1] != 1<<31-1 {
			return f0[n-1]
		}
		x := q1[0]
		step1 := f1[x] + 1
		if x+1 < n && step1 < f1[x+1] {
			f1[x+1] = step1
			q1 = append(q1, x+1)
		}
		if x-1 < n && step1 < f1[x-1] {
			f1[x-1] = step1
			q1 = append(q1, x-1)
		}
		for _, y := range h[arr[x]] {
			if step1 < f1[y] {
				f1[y] = step1
				q1 = append(q1, y)
			}
		}
		if f1[0] != 1<<31-1 {
			return f1[0]
		}
	}
	return -1
}
