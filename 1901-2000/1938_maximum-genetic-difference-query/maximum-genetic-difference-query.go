package main

// Tags:
// Trie
// DFS
// Sorting

func maxGeneticDifference(parents []int, queries [][]int) []int {
	var root int
	n := len(parents)
	g := make([][]int, n)
	for i, p := range parents {
		if p == -1 {
			root = i
		} else {
			g[p] = append(g[p], i)
		}
	}
	qs := make([][][]int, n)
	for i, q := range queries {
		qs[q[0]] = append(qs[q[0]], []int{i, q[1]})
	}

	out := make([]int, len(queries))

	dfs(&Trie{}, g, qs, out, root)

	return out
}

func dfs(t *Trie, g [][]int, qs [][][]int, out []int, node int) {
	t.Insert(node)
	for _, q := range qs[node] {
		out[q[0]] = t.Query(q[1])
	}
	for _, v := range g[node] {
		dfs(t, g, qs, out, v)
	}
	t.Delete(node)
}

type Trie struct {
	child [2]*Trie
	count int
}

func (t *Trie) Insert(x int) {
	for i := 31; i >= 0; i-- {
		idx := (x >> i) & 1
		if t.child[idx] == nil {
			t.child[idx] = &Trie{}
		}
		t = t.child[idx]
		t.count++
	}
}

func (t *Trie) Delete(x int) {
	for i := 31; i >= 0; i-- {
		idx := (x >> i) & 1
		t = t.child[idx]
		t.count--
	}
}

func (t *Trie) Query(val int) int {
	var out int
	for i := 31; i >= 0 && t != nil; i-- {
		idx := (val >> i) & 1
		if t.child[1^idx] != nil && t.child[1^idx].count > 0 {
			t = t.child[1^idx]
			out |= 1 << i
		} else {
			t = t.child[idx]
		}
	}
	return out
}
