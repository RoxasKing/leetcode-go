package main

import (
	"container/heap"
	"sort"
)

/*
  You are given a 2D integer array intervals, where intervals[i] = [lefti, righti] describes the ith interval starting at lefti and ending at righti (inclusive). The size of an interval is defined as the number of integers it contains, or more formally righti - lefti + 1.

  You are also given an integer array queries. The answer to the jth query is the size of the smallest interval i such that lefti <= queries[j] <= righti. If no such interval exists, the answer is -1.

  Return an array containing the answers to the queries.

  Example 1:
    Input: intervals = [[1,4],[2,4],[3,6],[4,4]], queries = [2,3,4,5]
    Output: [3,3,1,4]
    Explanation: The queries are processed as follows:
      - Query = 2: The interval [2,4] is the smallest interval containing 2. The answer is 4 - 2 + 1 = 3.
      - Query = 3: The interval [2,4] is the smallest interval containing 3. The answer is 4 - 2 + 1 = 3.
      - Query = 4: The interval [4,4] is the smallest interval containing 4. The answer is 4 - 4 + 1 = 1.
      - Query = 5: The interval [3,6] is the smallest interval containing 5. The answer is 6 - 3 + 1 = 4.

  Example 2:
    Input: intervals = [[2,3],[2,5],[1,8],[20,25]], queries = [2,19,5,22]
    Output: [2,-1,4,6]
    Explanation: The queries are processed as follows:
      - Query = 2: The interval [2,3] is the smallest interval containing 2. The answer is 3 - 2 + 1 = 2.
      - Query = 19: None of the intervals contain 19. The answer is -1.
      - Query = 5: The interval [2,5] is the smallest interval containing 5. The answer is 5 - 2 + 1 = 4.
      - Query = 22: The interval [20,25] is the smallest interval containing 22. The answer is 25 - 20 + 1 = 6.

  Constraints:
    1. 1 <= intervals.length <= 10^5
    2. 1 <= queries.length <= 10^5
    3. intervals[i].length == 2
    4. 1 <= lefti <= righti <= 10^7
    5. 1 <= queries[j] <= 10^7

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-interval-to-include-each-query
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Priority Queue(Heap Sort)
func minInterval(intervals [][]int, queries []int) []int {
	m, n := len(intervals), len(queries)

	idxs := make([]int, n)
	for i := range idxs {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool { return queries[idxs[i]] < queries[idxs[j]] })
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	idx := 0
	pq := MinHeap{}
	out := make([]int, n)

	for _, i := range idxs {
		query := queries[i]
		for ; idx < m && intervals[idx][0] <= query; idx++ {
			heap.Push(&pq, &interval{length: intervals[idx][1] - intervals[idx][0] + 1, right: intervals[idx][1]})
		}

		for pq.Len() > 0 && pq[0].right < query {
			heap.Pop(&pq)
		}

		if pq.Len() > 0 {
			out[i] = pq[0].length
		} else {
			out[i] = -1
		}
	}

	return out
}

type interval struct {
	length, right int
}

type MinHeap []*interval

func (pq MinHeap) Len() int            { return len(pq) }
func (pq MinHeap) Less(i, j int) bool  { return pq[i].length < pq[j].length }
func (pq MinHeap) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *MinHeap) Push(x interface{}) { *pq = append(*pq, x.(*interval)) }
func (pq *MinHeap) Pop() interface{} {
	top := pq.Len() - 1
	out := (*pq)[top]
	*pq = (*pq)[:top]
	return out
}
