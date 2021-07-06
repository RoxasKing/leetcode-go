package main

import (
	"sort"
	"strconv"
)

// Tags:
// Hash

func displayTable(orders [][]string) [][]string {
	names := map[string]bool{}
	tnos := []int{}
	counts := map[int]map[string]int{}
	for _, order := range orders {
		table, name := order[1], order[2]
		names[name] = true
		tno, _ := strconv.Atoi(table)
		if counts[tno] == nil {
			tnos = append(tnos, tno)
			counts[tno] = map[string]int{}
		}
		counts[tno][name]++
	}
	out := make([][]string, 0, len(counts)+1)
	topics := make([]string, 0, len(names)+1)
	topics = append(topics, "Table")
	for name := range names {
		topics = append(topics, name)
	}
	sort.Strings(topics[1:])
	out = append(out, topics)
	sort.Ints(tnos)
	for _, tno := range tnos {
		tmp := make([]string, 0, len(names)+1)
		tmp = append(tmp, strconv.Itoa(tno))
		for _, name := range topics[1:] {
			tmp = append(tmp, strconv.Itoa(counts[tno][name]))
		}
		out = append(out, tmp)
	}
	return out
}
