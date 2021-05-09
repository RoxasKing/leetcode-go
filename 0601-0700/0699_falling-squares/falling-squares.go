package main

import (
	"sort"
)

/*
  There are several squares being dropped onto the X-axis of a 2D plane.

  You are given a 2D integer array positions where positions[i] = [lefti, sideLengthi] represents the ith square with a side length of sideLengthi that is dropped with its left edge aligned with X-coordinate lefti.

  Each square is dropped one at a time from a height above any landed squares. It then falls downward (negative Y direction) until it either lands on the top side of another square or on the X-axis. A square brushing the left/right side of another square does not count as landing on it. Once it lands, it freezes in place and cannot be moved.

  After each square is dropped, you must record the height of the current tallest stack of squares.

  Return an integer array ans where ans[i] represents the height described above after dropping the ith square.

  Example 1:
    Input: positions = [[1,2],[2,3],[6,1]]
    Output: [2,5,5]
    Explanation:
      After the first drop, the tallest stack is square 1 with a height of 2.
      After the second drop, the tallest stack is squares 1 and 2 with a height of 5.
      After the third drop, the tallest stack is still squares 1 and 2 with a height of 5.
      Thus, we return an answer of [2, 5, 5].

  Example 2:
    Input: positions = [[100,100],[200,100]]
    Output: [100,100]
    Explanation:
      After the first drop, the tallest stack is square 1 with a height of 100.
      After the second drop, the tallest stack is either square 1 or square 2, both with heights of 100.
      Thus, we return an answer of [100, 100].
      Note that square 2 only brushes the right side of square 1, which does not count as landing on it.

  Constraints:
    1. 1 <= positions.length <= 1000
    2. 1 <= lefti <= 10^8
    3. 1 <= sideLengthi <= 10^6

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/falling-squares
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Segment Tree
func fallingSquares(positions [][]int) []int {
	set := make(map[int]struct{})
	for _, pos := range positions {
		set[pos[0]] = struct{}{}
		set[pos[0]+pos[1]-1] = struct{}{}
	}
	idxs := []int{}
	for pos := range set {
		idxs = append(idxs, pos)
	}
	sort.Ints(idxs)
	dict := make(map[int]int)
	for i, pos := range idxs {
		dict[pos] = i
	}
	st := NewSegmentTree(len(idxs))

	n := len(positions)
	out := make([]int, n)
	for i, pos := range positions {
		l, r := dict[pos[0]], dict[pos[0]+pos[1]-1]
		base := st.Query(1, st.s, st.t, l, r)
		st.Update(1, st.s, st.t, l, r, base+pos[1])
		out[i] = st.Query(1, st.s, st.t, st.s, st.t)
	}
	return out
}

type SegmentTree struct {
	f, b []int
	s, t int
}

func NewSegmentTree(n int) SegmentTree {
	f := make([]int, 4*n)
	b := make([]int, 4*n)
	s, t := 0, n-1
	return SegmentTree{
		f: f,
		b: b,
		s: s,
		t: t,
	}
}

func (st *SegmentTree) Query(idx, s, t, l, r int) int {
	if t < l || s > r {
		return 0
	}
	if l <= s && t <= r {
		return st.f[idx]
	}
	m := (s + t) >> 1
	if st.b[idx] != 0 {
		st.f[idx<<1] = st.b[idx]
		st.f[idx<<1+1] = st.b[idx]
		st.b[idx<<1] = st.b[idx]
		st.b[idx<<1+1] = st.b[idx]
		st.b[idx] = 0
	}
	return Max(st.Query(idx<<1, s, m, l, r), st.Query(idx<<1+1, m+1, t, l, r))
}

func (st *SegmentTree) Update(idx, s, t, l, r, val int) {
	if t < l || s > r {
		return
	}
	if l <= s && t <= r {
		st.f[idx] = val
		st.b[idx] = val
		return
	}
	m := (s + t) >> 1
	if st.b[idx] != 0 && s != t {
		st.f[idx<<1] = st.b[idx]
		st.f[idx<<1+1] = st.b[idx]
		st.b[idx<<1] = st.b[idx]
		st.b[idx<<1+1] = st.b[idx]
		st.b[idx] = 0
	}
	st.Update(idx<<1, s, m, l, r, val)
	st.Update(idx<<1+1, m+1, t, l, r, val)
	st.f[idx] = Max(st.f[idx], Max(st.f[idx<<1], st.f[idx<<1+1]))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
