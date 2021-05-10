package main

import (
	"sort"
)

/*
  A city's skyline is the outer contour of the silhouette formed by all the buildings in that city when viewed from a distance. Given the locations and heights of all the buildings, return the skyline formed by these buildings collectively.

  The geometric information of each building is given in the array buildings where buildings[i] = [lefti, righti, heighti]:
    1. lefti is the x coordinate of the left edge of the ith building.
    2. righti is the x coordinate of the right edge of the ith building.
    3. heighti is the height of the ith building.

  You may assume all buildings are perfect rectangles grounded on an absolutely flat surface at height 0.

  The skyline should be represented as a list of "key points" sorted by their x-coordinate in the form [[x1,y1],[x2,y2],...]. Each key point is the left endpoint of some horizontal segment in the skyline except the last point in the list, which always has a y-coordinate 0 and is used to mark the skyline's termination where the rightmost building ends. Any ground between the leftmost and rightmost buildings should be part of the skyline's contour.

  Note: There must be no consecutive horizontal lines of equal height in the output skyline. For instance, [...,[2 3],[4 5],[7 5],[11 5],[12 7],...] is not acceptable; the three lines of height 5 should be merged into one in the final output as such: [...,[2 3],[4 5],[12 7],...]

  Example 1:
    Input: buildings = [[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]
    Output: [[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]
    Explanation:
      Figure A shows the buildings of the input.
      Figure B shows the skyline formed by those buildings. The red points in figure B represent the key points in the output list.

  Example 2:
    Input: buildings = [[0,2,3],[2,5,3]]
    Output: [[0,3],[5,0]]

  Constraints:
    1. 1 <= buildings.length <= 10^4
    2. 0 <= lefti < righti <= 2^31 - 1
    3. 1 <= heighti <= 2^31 - 1
    4. buildings is sorted by lefti in non-decreasing order.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/the-skyline-problem
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Segment Tree
func getSkyline(buildings [][]int) [][]int {
	set := make(map[int]struct{})
	for _, building := range buildings {
		l, r := building[0], building[1]
		set[l] = struct{}{}
		set[r] = struct{}{}
	}
	arr := make([]int, 0, len(set))
	for building := range set {
		arr = append(arr, building)
	}
	sort.Ints(arr)
	dict := make(map[int]int)
	for i, building := range arr {
		dict[building] = i
	}
	n := len(arr)

	st := NewSegmentTree(n)
	for _, building := range buildings {
		l, r, val := dict[building[0]], dict[building[1]]-1, building[2]
		st.Update(1, st.s, st.t, l, r, val)
	}

	out := [][]int{}
	for _, b := range arr {
		idx := dict[b]
		h := st.Query(1, st.s, st.t, idx, idx)
		if len(out) == 0 || out[len(out)-1][1] != h {
			out = append(out, []int{b, h})
		}
	}
	return out
}

type SegmentTree struct {
	f, b []int
	s, t int
}

func NewSegmentTree(n int) SegmentTree {
	return SegmentTree{
		f: make([]int, 4*n),
		b: make([]int, 4*n),
		s: 0,
		t: n - 1,
	}
}

func (st *SegmentTree) Query(idx, s, t, l, r int) int {
	if t < l || s > r {
		return 0
	}
	if l <= s && t <= r {
		if st.b[idx] != 0 {
			st.f[idx] = Max(st.f[idx], st.b[idx])
			st.b[idx] = 0
		}
		return st.f[idx]
	}
	m := (s + t) >> 1
	if st.b[idx] != 0 && s != t {
		st.f[idx<<1] = Max(st.f[idx<<1], st.b[idx])
		st.f[idx<<1+1] = Max(st.f[idx<<1+1], st.b[idx])
		st.b[idx<<1] = Max(st.b[idx<<1], st.b[idx])
		st.b[idx<<1+1] = Max(st.b[idx<<1+1], st.b[idx])
		st.b[idx] = 0
	}
	return Max(st.Query(idx<<1, s, m, l, r), st.Query(idx<<1+1, m+1, t, l, r))
}

func (st *SegmentTree) Update(idx, s, t, l, r, val int) {
	if t < l || s > r {
		return
	}
	if l <= s && t <= r {
		if st.b[idx] != 0 {
			st.f[idx] = Max(st.f[idx], st.b[idx])
		}
		st.f[idx] = Max(st.f[idx], val)
		st.b[idx] = st.f[idx]
		return
	}
	m := (s + t) >> 1
	if st.b[idx] != 0 && s != t {
		st.f[idx<<1] = Max(st.f[idx<<1], st.b[idx])
		st.f[idx<<1+1] = Max(st.f[idx<<1+1], st.b[idx])
		st.b[idx<<1] = st.f[idx<<1]
		st.b[idx<<1+1] = st.f[idx<<1+1]
		st.b[idx] = 0
	}
	st.Update(idx<<1, s, m, l, r, val)
	st.Update(idx<<1+1, m+1, t, l, r, val)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
