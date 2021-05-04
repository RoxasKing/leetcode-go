package main

import (
	"container/heap"
)

/*
  It is a sweltering summer day, and a boy wants to buy some ice cream bars.

  At the store, there are n ice cream bars. You are given an array costs of length n, where costs[i] is the price of the ith ice cream bar in coins. The boy initially has coins coins to spend, and he wants to buy as many ice cream bars as possible.

  Return the maximum number of ice cream bars the boy can buy with coins coins.

  Note: The boy can buy the ice cream bars in any order.

  Example 1:
    Input: costs = [1,3,2,4,1], coins = 7
    Output: 4
    Explanation: The boy can buy ice cream bars at indices 0,1,2,4 for a total price of 1 + 3 + 2 + 1 = 7.

  Example 2:
    Input: costs = [10,6,8,7,7,8], coins = 5
    Output: 0
    Explanation: The boy cannot afford any of the ice cream bars.

  Example 3:
    Input: costs = [1,6,3,1,2,5], coins = 20
    Output: 6
    Explanation: The boy can buy all the ice cream bars for a total price of 1 + 6 + 3 + 1 + 2 + 5 = 18.

  Constraints:
    1. costs.length == n
    2. 1 <= n <= 10^5
    3. 1 <= costs[i] <= 10^5
    4. 1 <= coins <= 10^8

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-ice-cream-bars
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Priority Queue(Heap Sort)
func maxIceCream(costs []int, coins int) int {
	pq := PriorityQueue{}
	for _, c := range costs {
		heap.Push(&pq, c)
	}
	out := 0
	for pq.Len() > 0 && coins-pq[0] >= 0 {
		coins -= heap.Pop(&pq).(int)
		out++
	}
	return out
}

type PriorityQueue []int

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i] < pq[j] }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(int)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := pq.Len() - 1
	out := (*pq)[top]
	*pq = (*pq)[:top]
	return out
}
