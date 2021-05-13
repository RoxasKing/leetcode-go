package main

import (
	"math/rand"
	"sort"
)

/*
  Given a list of non-overlapping axis-aligned rectangles rects, write a function pick which randomly and uniformily picks an integer point in the space covered by the rectangles.

  Note:
    1. An integer point is a point that has integer coordinates.
    2. A point on the perimeter of a rectangle is included in the space covered by the rectangles.
    3. ith rectangle = rects[i] = [x1,y1,x2,y2], where [x1, y1] are the integer coordinates of the bottom-left corner, and [x2, y2] are the integer coordinates of the top-right corner.
    4. length and width of each rectangle does not exceed 2000.
    5. 1 <= rects.length <= 100
    6. pick return a point as an array of integer coordinates [p_x, p_y]
    7. pick is called at most 10000 times.

  Example 1:
    Input:
      ["Solution","pick","pick","pick"]
      [[[[1,1,5,5]]],[],[],[]]
    Output:
      [null,[4,1],[4,1],[3,3]]

  Example 2:
    Input:
      ["Solution","pick","pick","pick","pick","pick"]
      [[[[-2,-2,-1,-1],[1,0,3,0]]],[],[],[],[],[]]
    Output:
      [null,[-1,-2],[2,0],[-2,-1],[3,0],[-2,-2]]

  Explanation of Input Syntax:
    The input is two lists: the subroutines called and their arguments. Solution's constructor has one argument, the array of rectangles rects. pick has no arguments. Arguments are always wrapped with a list, even if there aren't any.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/random-point-in-non-overlapping-rectangles
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search

type Solution struct {
	n   int
	arr [][]int
}

func Constructor(rects [][]int) Solution {
	arr := [][]int{}
	start := 0
	for _, rect := range rects {
		a, b, c, d := rect[0], rect[1], rect[2], rect[3]
		arr = append(arr, []int{start, d - b + 1, a, b})
		start += (c - a + 1) * (d - b + 1)
	}
	return Solution{
		n:   start,
		arr: arr,
	}
}

func (this *Solution) Pick() []int {
	idx := rand.Intn(this.n)
	i := sort.Search(len(this.arr), func(i int) bool { return this.arr[i][0] > idx }) - 1
	extra := idx - this.arr[i][0]
	n, x, y := this.arr[i][1], this.arr[i][2], this.arr[i][3]
	return []int{x + extra/n, y + extra%n}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */
