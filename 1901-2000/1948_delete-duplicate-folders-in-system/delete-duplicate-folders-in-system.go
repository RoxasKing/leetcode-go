package main

import (
	"sort"
	"strings"
)

// Tags:
// Trie
// Hash

func deleteDuplicateFolder(paths [][]string) [][]string {
	t := &trie{}
	for _, path := range paths {
		t.insert(path)
	}

	seen := map[string]*trie{}
	dedupe(t, seen)

	out := make([][]string, 0, len(paths))
	for _, next := range t.child {
		dfs_out(next, []string{}, &out)
	}
	return out
}

func dedupe(t *trie, seen map[string]*trie) string {
	subfolders := make([]string, 0, len(t.child))
	for _, next := range t.child {
		subfolders = append(subfolders, dedupe(next, seen))
	}
	sort.Strings(subfolders)
	sub := strings.Join(subfolders, "")
	if len(subfolders) > 0 {
		if seen[sub] == nil {
			seen[sub] = t
		} else {
			seen[sub].isDel = true
			t.isDel = true
		}
	}
	return "(" + t.fName + sub + ")"
}

func dfs_out(t *trie, path []string, out *[][]string) {
	if t.isDel {
		return
	}
	path = append(path, t.fName)
	*out = append(*out, append(make([]string, 0, len(path)), path...))
	for _, next := range t.child {
		dfs_out(next, path, out)
	}
}

type trie struct {
	child map[string]*trie
	fName string
	isDel bool
}

func (t *trie) insert(path []string) {
	for _, p := range path {
		if t.child == nil {
			t.child = map[string]*trie{}
		}
		if t.child[p] == nil {
			t.child[p] = &trie{}
		}
		t = t.child[p]
		t.fName = p
	}
}
