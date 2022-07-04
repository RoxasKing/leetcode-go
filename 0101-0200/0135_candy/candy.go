package main

// Difficulty:
// Hard

// Tags:
// Greedy

func candy(ratings []int) int {
	n := len(ratings)
	f := make([]int, n)
	q := []int{}
	for i, v := range ratings {
		if i-1 >= 0 && ratings[i-1] < v || i+1 < n && v > ratings[i+1] {
			continue
		}
		f[i] = 1
		q = append(q, i)
	}
	for ; len(q) > 0; q = q[1:] {
		i := q[0]
		if i-1 >= 0 && ratings[i-1] > ratings[i] && f[i-1] < f[i]+1 {
			f[i-1] = f[i] + 1
			q = append(q, i-1)
		}
		if i+1 < n && ratings[i+1] > ratings[i] && f[i+1] < f[i]+1 {
			f[i+1] = f[i] + 1
			q = append(q, i+1)
		}
	}
	o := 0
	for _, x := range f {
		o += x
	}
	return o
}
