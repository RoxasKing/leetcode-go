package main

import "container/heap"

/*
  A super ugly number is a positive integer whose prime factors are in the array primes.

  Given an integer n and an array of integers primes, return the nth super ugly number.

  The nth super ugly number is guaranteed to fit in a 32-bit signed integer.

  Example 1:
    Input: n = 12, primes = [2,7,13,19]
    Output: 32
    Explanation: [1,2,4,7,8,13,14,16,19,26,28,32] is the sequence of the first 12 super ugly numbers given primes = [2,7,13,19].

  Example 2:
    Input: n = 1, primes = [2,3,5]
    Output: 1
    Explanation: 1 has no prime factors, therefore all of its prime factors are in the array primes = [2,3,5].

  Constraints:
    1. 1 <= n <= 10^6
    2. 1 <= primes.length <= 100
    3. 2 <= primes[i] <= 1000
    4. primes[i] is guaranteed to be a prime number.
    5. All the values of primes are unique and sorted in ascending order.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/super-ugly-number
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
// Priority Queue

func nthSuperUglyNumber(n int, primes []int) int {
	h := MinHeap{}
	for _, prime := range primes {
		heap.Push(&h, &Pointer{idx: 0, mul: prime})
	}
	for i := 1; i < n; i++ {
		for f[h[0].idx]*h[0].mul <= f[i-1] {
			p := heap.Pop(&h).(*Pointer)
			heap.Push(&h, &Pointer{idx: p.idx + 1, mul: p.mul})
		}
		f[i] = f[h[0].idx] * h[0].mul
	}
	return f[n-1]
}

var f = [100000]int{1}

type Pointer struct {
	idx int
	mul int
}

type MinHeap []*Pointer

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return f[h[i].idx]*h[i].mul < f[h[j].idx]*h[j].mul }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*Pointer)) }
func (h *MinHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
