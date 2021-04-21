package main

import (
	"sort"
)

/*
  You are given a string s, and an array of pairs of indices in the string pairs where pairs[i] = [a, b] indicates 2 indices(0-indexed) of the string.
  You can swap the characters at any pair of indices in the given pairs any number of times.
  Return the lexicographically smallest string that s can be changed to after using the swaps.

  Example 1:
    Input: s = "dcab", pairs = [[0,3],[1,2]]
    Output: "bacd"
    Explaination:
    Swap s[0] and s[3], s = "bcad"
    Swap s[1] and s[2], s = "bacd"

  Example 2:
    Input: s = "dcab", pairs = [[0,3],[1,2],[0,2]]
    Output: "abcd"
    Explaination:
    Swap s[0] and s[3], s = "bcad"
    Swap s[0] and s[2], s = "acbd"
    Swap s[1] and s[2], s = "abcd"

  Example 3:
    Input: s = "cba", pairs = [[0,1],[1,2]]
    Output: "abc"
    Explaination:
    Swap s[0] and s[1], s = "bca"
    Swap s[1] and s[2], s = "bac"
    Swap s[0] and s[1], s = "abc"

  Constraints:
    1 <= s.length <= 10^5
    0 <= pairs.length <= 10^5
    0 <= pairs[i][0], pairs[i][1] < s.length
    s only contains lower case English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/smallest-string-with-swaps
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Union-Find
func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)

	uf := newUnionFind(n)
	for _, pair := range pairs {
		uf.union(pair[0], pair[1])
	}

	groups := make(map[int][]int)
	for i := 0; i < n; i++ {
		a := uf.find(i)
		groups[a] = append(groups[a], i)
	}

	bs := []byte(s)
	for _, group := range groups {
		if len(group) < 2 {
			continue
		}
		arr := make([]byte, len(group))
		for i := range group {
			arr[i] = bs[group[i]]
		}
		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
		for i := range arr {
			bs[group[i]] = arr[i]
		}
	}

	return string(bs)
}

type unionFind struct {
	ancestor []int
	isEnd    []bool
}

func newUnionFind(n int) unionFind {
	ancestor := make([]int, n)
	isEnd := make([]bool, n)
	for i := range ancestor {
		ancestor[i] = i
	}
	return unionFind{ancestor: ancestor, isEnd: isEnd}
}

func (uf unionFind) find(x int) int {
	if uf.isEnd[uf.ancestor[x]] {
		return uf.ancestor[x]
	}

	if uf.ancestor[x] != x {
		uf.ancestor[x] = uf.find(uf.ancestor[x])
	}
	uf.isEnd[x] = false
	uf.isEnd[uf.ancestor[x]] = true

	return uf.ancestor[x]
}

func (uf unionFind) union(x, y int) {
	xAncestor := uf.find(x)
	yAncestor := uf.find(y)
	uf.ancestor[xAncestor] = yAncestor
	uf.isEnd[xAncestor] = false
	uf.isEnd[yAncestor] = true
}
