package main

/*
  You are given an even integer n​​​​​​. You initially have a permutation perm of size n​​ where perm[i] == i​ (0-indexed)​​​​.

  In one operation, you will create a new array arr, and for each i:

  If i % 2 == 0, then arr[i] = perm[i / 2].
  If i % 2 == 1, then arr[i] = perm[n / 2 + (i - 1) / 2].
  You will then assign arr​​​​ to perm.

  Return the minimum non-zero number of operations you need to perform on perm to return the permutation to its initial value.

Example 1:
  Input: n = 2
  Output: 1
  Explanation: nums = [0,1] initially.
  After the 1st operation, nums = [0,1]
  So it takes only 1 operation.

Example 2:
  Input: n = 4
  Output: 2
  Explanation: nums = [0,1,2,3] initially.
  After the 1st operation, nums = [0,2,1,3]
  After the 2nd operation, nums = [0,1,2,3]
  So it takes only 2 operations.

Example 3:
  Input: n = 6
  Output: 4

Constraints:
  1. 2 <= n <= 1000
  2. n​​​​​​ is even.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-number-of-operations-to-reinitialize-a-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Union-Find
func reinitializePermutation(n int) int {
	uf := NewUnionFind(n)
	for i := 0; i < n; i++ {
		x := i / 2
		if i&1 == 1 {
			x = n/2 + (i-1)/2
		}
		uf.Union(i, x)
	}

	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		cnt[uf.Find(i)]++
	}

	out := 0
	for i := 0; i < n; i++ {
		out = Max(out, cnt[i])
	}
	return out
}

type unionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return unionFind{
		parent: parent,
		size:   size,
	}
}

func (uf unionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf unionFind) Union(x, y int) {
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

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
