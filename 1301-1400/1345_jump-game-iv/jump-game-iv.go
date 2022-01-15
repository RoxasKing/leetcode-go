package main

// Difficulty:
// Hard

// Tags:
// Hash
// BFS

func minJumps(arr []int) int {
	n := len(arr)
	f1 := make([]int, n)
	f2 := make([]int, n)
	for i := range f1 {
		f1[i] = 1<<31 - 1
		f2[i] = 1<<31 - 1
	}
	f1[0], f2[n-1] = 0, 0
	h := map[int][]int{}
	for i, num := range arr {
		h[num] = append(h[num], i)
	}
	for q1, q2 := []int{0}, []int{n - 1}; len(q1) > 0; q1, q2 = q1[1:], q2[1:] {
		a := q1[0]
		step1 := f1[a] + 1
		if a+1 < n && step1 < f1[a+1] {
			f1[a+1] = step1
			q1 = append(q1, a+1)
		}
		if a-1 >= 0 && step1 < f1[a-1] {
			f1[a-1] = step1
			q1 = append(q1, a-1)
		}
		for _, b := range h[arr[a]] {
			if step1 < f1[b] {
				f1[b] = step1
				q1 = append(q1, b)
			}
		}
		if f1[n-1] != 1<<31-1 {
			return f1[n-1]
		}
		x := q2[0]
		step2 := f2[x] + 1
		if x+1 < n && step2 < f2[x+1] {
			f2[x+1] = step2
			q2 = append(q2, x+1)
		}
		if x-1 >= 0 && step2 < f2[x-1] {
			f2[x-1] = step2
			q2 = append(q2, x-1)
		}
		for _, y := range h[arr[x]] {
			if step2 < f2[y] {
				f2[y] = step2
				q2 = append(q2, y)
			}
		}
		if f2[0] != 1<<31-1 {
			return f2[0]
		}
	}
	return -1
}
