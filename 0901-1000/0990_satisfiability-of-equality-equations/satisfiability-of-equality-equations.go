package main

/*
  Given an array equations of strings that represent relationships between variables, each string equations[i] has length 4 and takes one of two different forms: "a==b" or "a!=b".  Here, a and b are lowercase letters (not necessarily different) that represent one-letter variable names.

  Return true if and only if it is possible to assign integers to variable names so as to satisfy all the given equations.

  Example 1:
    Input: ["a==b","b!=a"]
    Output: false
    Explanation: If we assign say, a = 1 and b = 1, then the first equation is satisfied, but not the second.  There is no way to assign the variables to satisfy both equations.

  Example 2:
    Input: ["b==a","a==b"]
    Output: true
    Explanation: We could assign a = 1 and b = 1 to satisfy both equations.

  Example 3:
    Input: ["a==b","b==c","a==c"]
    Output: true

  Example 4:
    Input: ["a==b","b!=c","c==a"]
    Output: false

  Example 5:
    Input: ["c==c","b==d","x!=z"]
    Output: true

  Note:
    1. 1 <= equations.length <= 500
    2. equations[i].length == 4
    3. equations[i][0] and equations[i][3] are lowercase letters
    4. equations[i][1] is either '=' or '!'
    5. equations[i][2] is '='

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/satisfiability-of-equality-equations
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Union-Find
func equationsPossible(equations []string) bool {
	uf := NewUnionFind(26)
	for _, e := range equations {
		x, y := int(e[0]-'a'), int(e[3]-'a')
		if e[1] == '=' {
			uf.Union(x, y)
		}
	}
	for _, e := range equations {
		x, y := int(e[0]-'a'), int(e[3]-'a')
		x, y = uf.Find(x), uf.Find(y)
		if e[1] == '!' && x == y {
			return false
		}
	}
	return true
}

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return UnionFind{
		parent: parent,
		size:   size,
	}
}

func (uf UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf UnionFind) Union(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}
	if uf.size[x] < uf.size[y] {
		x, y = y, x
	}
	uf.parent[y] = x
	uf.size[x] += uf.size[y]
}
