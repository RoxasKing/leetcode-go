package main

import (
	"math/rand"
	"sort"
)

/*
  Given a blacklist blacklist containing unique integers from [0, n), write a function to return a uniform random integer from [0, n) which is NOT in blacklist.

  Optimize it such that it minimizes the call to system’s Math.random().

  Note:
    1. 1 <= n <= 1000000000
    2. 0 <= blacklist.length < min(100000, n)
    3. [0, n) does NOT include n. See interval notation.

  Example 1:
    Input:
      ["Solution","pick","pick","pick"]
      [[1,[]],[],[],[]]
    Output: [null,0,0,0]

  Example 2:
    Input:
      ["Solution","pick","pick","pick"]
      [[2,[]],[],[],[]]
    Output: [null,1,1,1]

  Example 3:
    Input:
      ["Solution","pick","pick","pick"]
      [[3,[1]],[],[],[]]
    Output: [null,0,0,2]

  Example 4:
    Input:
      ["Solution","pick","pick","pick"]
      [[4,[2]],[],[],[]]
    Output: [null,1,3,1]

  Explanation of Input Syntax:
    The input is two lists: the subroutines called and their arguments. Solution's constructor has two arguments, n and the blacklist blacklist. pick has no arguments. Arguments are always wrapped with a list, even if there aren't any.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/random-pick-with-blacklist
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Binary Search

type Solution struct {
	n int
	b []int
}

func Constructor(n int, blacklist []int) Solution {
	sort.Ints(blacklist)
	return Solution{
		n: n,
		b: blacklist,
	}
}

func (this *Solution) Pick() int {
	x := rand.Intn(this.n - len(this.b))
	l, r := 0, len(this.b)-1
	for l < r {
		m := (l + r + 1) >> 1
		if this.b[m]-m <= x {
			l = m
		} else {
			r = m - 1
		}
	}
	if l == r && this.b[l]-l <= x {
		return x + l + 1
	}
	return x
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
