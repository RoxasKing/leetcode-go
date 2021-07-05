package main

import "strconv"

// BFS + Hash

func openLock(deadends []string, target string) int {
	dead := [10000]bool{}
	for _, d := range deadends {
		num, _ := strconv.Atoi(d)
		dead[num] = true
	}
	walked := [10000]bool{0: true}
	targetNum, _ := strconv.Atoi(target)
	out := -1
	q := [][2]int{{0, 0}}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		num, step := x[0], x[1]
		if dead[num] {
			continue
		}
		if num == targetNum {
			out = step
			break
		}
		step++
		for base := 1000; base >= 1; base /= 10 {
			i := (num / base) % 10
			next := num - i*base + (i+1)%10*base
			if !walked[next] {
				walked[next] = true
				q = append(q, [2]int{next, step})
			}
			next = num - i*base + (i+9)%10*base
			if !walked[next] {
				walked[next] = true
				q = append(q, [2]int{next, step})
			}
		}
	}
	return out
}
