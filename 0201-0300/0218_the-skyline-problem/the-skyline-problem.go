package main

import (
	"container/heap"
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

// Priority Queue(Heap Sort)
func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	points := make([]int, 0, n<<1)
	for _, b := range buildings {
		points = append(points, b[0], b[1])
	}
	sort.Ints(points)
	h := MaxHeap{}
	heap.Push(&h, &buildInf{rBound: 1<<63 - 1, height: 0})
	idx := 0
	out := [][]int{}

	for _, p := range points {
		for ; idx < n && buildings[idx][0] == p; idx++ {
			heap.Push(&h, &buildInf{rBound: buildings[idx][1], height: buildings[idx][2]})
		}
		for h[0].rBound <= p {
			heap.Pop(&h)
		}
		if len(out) == 0 || out[len(out)-1][1] != h[0].height {
			out = append(out, []int{p, h[0].height})
		}
	}

	return out
}

type buildInf struct {
	rBound, height int
}

type MaxHeap []*buildInf

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].height > h[j].height }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(*buildInf)) }
func (h *MaxHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
