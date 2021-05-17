package main

/*
  Given an array arr that represents a permutation of numbers from 1 to n. You have a binary string of size n that initially has all its bits set to zero.

  At each step i (assuming both the binary string and arr are 1-indexed) from 1 to n, the bit at position arr[i] is set to 1. You are given an integer m and you need to find the latest step at which there exists a group of ones of length m. A group of ones is a contiguous substring of 1s such that it cannot be extended in either direction.

  Return the latest step at which there exists a group of ones of length exactly m. If no such group exists, return -1.

  Example 1:
    Input: arr = [3,5,1,2,4], m = 1
    Output: 4
    Explanation:
      Step 1: "00100", groups: ["1"]
      Step 2: "00101", groups: ["1", "1"]
      Step 3: "10101", groups: ["1", "1", "1"]
      Step 4: "11101", groups: ["111", "1"]
      Step 5: "11111", groups: ["11111"]
      The latest step at which there exists a group of size 1 is step 4.

  Example 2:
    Input: arr = [3,1,5,4,2], m = 2
    Output: -1
    Explanation:
      Step 1: "00100", groups: ["1"]
      Step 2: "10100", groups: ["1", "1"]
      Step 3: "10101", groups: ["1", "1", "1"]
      Step 4: "10111", groups: ["1", "111"]
      Step 5: "11111", groups: ["11111"]
      No group of size 2 exists during any step.

  Example 3:
    Input: arr = [1], m = 1
    Output: 1

  Example 4:
    Input: arr = [2,1], m = 2
    Output: 2

  Constraints:
    1. n == arr.length
    2. 1 <= n <= 10^5
    3. 1 <= arr[i] <= n
    4. All integers in arr are distinct.
    5. 1 <= m <= arr.length

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-latest-group-of-size-m
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Union-Find

func findLatestStep(arr []int, m int) int {
	n := len(arr)
	uf := NewUnionFind(n + 1)
	on := make([]bool, n+1)
	cnt := map[int]int{}

	out := -1
	for j, i := range arr {
		on[i] = true
		cnt[uf.GetSize(i)]++
		if i-1 > 0 && on[i-1] {
			cnt[uf.GetSize(i)]--
			cnt[uf.GetSize(i-1)]--
			uf.Union(i, i-1)
			cnt[uf.GetSize(i)]++
		}
		if i+1 <= n && on[i+1] {
			cnt[uf.GetSize(i)]--
			cnt[uf.GetSize(i+1)]--
			uf.Union(i, i+1)
			cnt[uf.GetSize(i)]++
		}
		if cnt[m] > 0 {
			out = j + 1
		}
	}
	return out
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
	return UnionFind{parent: parent, size: size}
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

func (uf UnionFind) GetSize(x int) int {
	return uf.size[uf.Find(x)]
}
