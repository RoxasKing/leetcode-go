package main

// Tags:
// Tags:
// Greedy Algorithm

func eliminateMaximum(dist []int, speed []int) int {
	cnt := make([]int, 1e5+1)
	for i := range dist {
		if dist[i]%speed[i] == 0 {
			cnt[dist[i]/speed[i]-1]++
		} else {
			cnt[dist[i]/speed[i]]++
		}
	}

	out := 0
	for i := range cnt {
		if out+cnt[i] > i+1 {
			out = i + 1
			break
		}
		out += cnt[i]
	}
	return out
}
