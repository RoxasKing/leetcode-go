package main

import (
	"sort"
)

// Tags:
// Backtracking + Hash + Graph
func findItinerary(tickets [][]string) []string {
	dict := make(map[string][]string)
	for _, ticket := range tickets {
		dict[ticket[0]] = append(dict[ticket[0]], ticket[1])
	}
	for k := range dict {
		sort.Strings(dict[k])
	}
	out := make([]string, 0, len(tickets)+1)
	out = append(out, "JFK")
	var backtrack func(string) bool
	backtrack = func(src string) bool {
		if _, ok := dict[src]; !ok || len(out) == len(tickets)+1 {
			return len(out) == len(tickets)+1
		}
		for _, dst := range dict[src] {
			out = append(out, dst)
			dict[src] = dict[src][1:]
			if backtrack(dst) {
				return true
			}
			out = out[:len(out)-1]
			dict[src] = append(dict[src], dst)
		}
		return false
	}
	_ = backtrack("JFK")
	return out
}

// DFS + Eulerian path + Hash + Graph
func findItinerary2(tickets [][]string) []string {
	dict := make(map[string][]string)
	for _, ticket := range tickets {
		dict[ticket[0]] = append(dict[ticket[0]], ticket[1])
	}
	for k := range dict {
		sort.Strings(dict[k])
	}
	out := make([]string, 0, len(tickets)+1)
	var dfs func(string)
	dfs = func(src string) {
		for len(dict[src]) != 0 {
			dst := dict[src][0]
			dict[src] = dict[src][1:]
			dfs(dst)
		}
		out = append(out, src)
	}
	dfs("JFK")
	for i := 0; i < len(out)>>1; i++ {
		out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]
	}
	return out
}
