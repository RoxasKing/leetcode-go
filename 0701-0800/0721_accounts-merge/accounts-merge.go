package main

import "sort"

// Union-Find
func accountsMerge(accounts [][]string) [][]string {
	n := len(accounts)

	// union all related accounts
	uf := newUnionFind(n)

	// record all emails account index
	emails := make(map[string]int)

	for i, account := range accounts {
		for _, email := range account[1:] {
			if idx, ok := emails[email]; ok {
				uf.union(idx, i)
			} else {
				emails[email] = i
			}
		}
	}

	accountsMap := make(map[int][]string)

	for i := range accounts {
		idx := uf.find(i)
		if _, ok := accountsMap[idx]; ok {
			continue
		}
		name := accounts[i][0]
		accountsMap[idx] = append(accountsMap[idx], name)
	}

	for email, i := range emails {
		idx := uf.find(i)
		accountsMap[idx] = append(accountsMap[idx], email)
	}

	out := make([][]string, 0, len(accountsMap))
	for _, account := range accountsMap {
		sort.Strings(account)
		out = append(out, account)
	}

	return out
}

type unionFind struct {
	ancestor []int
	isEnd    []bool
}

func newUnionFind(n int) unionFind {
	ancestor := make([]int, n)
	for i := range ancestor {
		ancestor[i] = i
	}
	isEnd := make([]bool, n)
	return unionFind{ancestor: ancestor, isEnd: isEnd}
}

func (uf unionFind) find(x int) int {
	if uf.isEnd[uf.ancestor[x]] {
		return uf.ancestor[x]
	}
	if uf.ancestor[x] != x {
		uf.ancestor[x] = uf.find(uf.ancestor[x])
		uf.isEnd[x] = false
		uf.isEnd[uf.ancestor[x]] = true
	}
	return uf.ancestor[x]
}

func (uf unionFind) union(x, y int) {
	ancestorX := uf.find(x)
	ancestorY := uf.find(y)
	uf.ancestor[ancestorX] = ancestorY
	uf.isEnd[ancestorX] = false
	uf.isEnd[ancestorY] = true
}
