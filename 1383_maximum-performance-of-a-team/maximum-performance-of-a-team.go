package main

import (
	"container/heap"
	"sort"
)

/*
  公司有编号为 1 到 n 的 n 个工程师，给你两个数组 speed 和 efficiency ，其中 speed[i] 和 efficiency[i] 分别代表第 i 位工程师的速度和效率。请你返回由最多 k 个工程师组成的 ​​​​​​最大团队表现值 ，由于答案可能很大，请你返回结果对 10^9 + 7 取余后的结果。

  团队表现值 的定义为：一个团队中「所有工程师速度的和」乘以他们「效率值中的最小值」。

  提示：
    1 <= n <= 10^5
    speed.length == n
    efficiency.length == n
    1 <= speed[i] <= 10^5
    1 <= efficiency[i] <= 10^8
    1 <= k <= n

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-performance-of-a-team
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Stack
func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	indexes := make([]int, n)
	for i := range indexes {
		indexes[i] = i
	}
	sort.Slice(indexes, func(i, j int) bool {
		return efficiency[indexes[i]] > efficiency[indexes[j]]
	})
	ph := speedHeap{}
	heap.Init(&ph)
	var speedSum int
	var max int64
	for _, index := range indexes {
		if ph.Len() == k {
			speedSum -= heap.Pop(&ph).(int)
		}
		speedSum += speed[index]
		heap.Push(&ph, speed[index])
		max = Max(max, int64(speedSum)*int64(efficiency[index]))
	}
	return int(max % (1e9 + 7))
}

type speedHeap []int

func (h speedHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h speedHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h speedHeap) Len() int            { return len(h) }
func (h *speedHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *speedHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:h.Len()-1]
	return res
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
