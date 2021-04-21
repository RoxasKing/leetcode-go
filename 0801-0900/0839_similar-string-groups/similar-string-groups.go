package main

/*
  Two strings X and Y are similar if we can swap two letters (in different positions) of X, so that it equals Y. Also two strings X and Y are similar if they are equal.

  For example, "tars" and "rats" are similar (swapping at positions 0 and 2), and "rats" and "arts" are similar, but "star" is not similar to "tars", "rats", or "arts".

  Together, these form two connected groups by similarity: {"tars", "rats", "arts"} and {"star"}.  Notice that "tars" and "arts" are in the same group even though they are not similar.  Formally, each group is such that a word is in the group if and only if it is similar to at least one other word in the group.

  We are given a list strs of strings where every string in strs is an anagram of every other string in strs. How many groups are there?

  Example 1:
    Input: strs = ["tars","rats","arts","star"]
    Output: 2

  Example 2:
    Input: strs = ["omv","ovm"]
    Output: 1

  Constraints:
    1 <= strs.length <= 100
    1 <= strs[i].length <= 1000
    sum(strs[i].length) <= 2 * 104
    strs[i] consists of lowercase letters only.
    All words in strs have the same length and are anagrams of each other.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/similar-string-groups
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Union-Find
func numSimilarGroups(strs []string) int {
	n := len(strs)
	uf := newUnionFind(n)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if isSimilar(strs[i], strs[j]) {
				uf.union(i, j)
			}
		}
	}
	mark := make([]bool, n)
	out := 0
	for i := range strs {
		x := uf.find(i)
		if !mark[x] {
			out++
			mark[x] = true
		}
	}
	return out
}

func isSimilar(s1, s2 string) bool {
	count := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
			if count > 2 {
				return false
			}
		}
	}
	return true
}

type unionFind struct {
	parent []int
	size   []int
}

func newUnionFind(n int) unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return unionFind{parent: parent, size: size}
}

func (uf unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf unionFind) union(x, y int) {
	x = uf.find(x)
	y = uf.find(y)
	if x == y {
		return
	}
	if uf.size[x] < uf.size[y] {
		x, y = y, x
	}
	uf.parent[y] = x
	uf.size[x] += uf.size[y]
}
