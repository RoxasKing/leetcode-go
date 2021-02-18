package main

/*
  N couples sit in 2N seats arranged in a row and want to hold hands. We want to know the minimum number of swaps so that every couple is sitting side by side. A swap consists of choosing any two people, then they stand up and switch seats.

  The people and seats are represented by an integer from 0 to 2N-1, the couples are numbered in order, the first couple being (0, 1), the second couple being (2, 3), and so on with the last couple being (2N-2, 2N-1).

  The couples' initial seating is given by row[i] being the value of the person who is initially sitting in the i-th seat.

  Example 1:
    Input: row = [0, 2, 1, 3]
    Output: 1
    Explanation: We only need to swap the second (row[1]) and third (row[2]) person.

  Example 2:
    Input: row = [3, 2, 0, 1]
    Output: 0
    Explanation: All couples are already seated side by side.

  Note:
    1. len(row) is even and in the range of [4, 60].
    2. row is guaranteed to be a permutation of 0...len(row)-1.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/couples-holding-hands
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Union-Find
func minSwapsCouples(row []int) int {
	n := len(row) >> 1
	uf := newUnionFind(n)
	for i := 0; i < len(row); i += 2 {
		c1, c2 := row[i]>>1, row[i+1]>>1
		if c1 == c2 {
			continue
		}
		if uf.find(c1) == uf.find(c2) {
			continue
		}
		uf.union(c1, c2)
	}

	out := 0
	mark := make([]bool, n)
	for i := 0; i < n; i++ {
		x := uf.find(i)
		if mark[x] {
			out++
			continue
		}
		mark[x] = true
	}
	return out
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
