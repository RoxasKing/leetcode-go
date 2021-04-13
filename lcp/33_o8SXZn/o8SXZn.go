package main

import (
	"container/heap"
	"math"
)

/*
  给定 N 个无限容量且初始均空的水缸，每个水缸配有一个水桶用来打水，第 i 个水缸配备的水桶容量记作 bucket[i]。小扣有以下两种操作：
    1. 升级水桶：选择任意一个水桶，使其容量增加为 bucket[i]+1
    2. 蓄水：将全部水桶接满水，倒入各自对应的水缸
  每个水缸对应最低蓄水量记作 vat[i]，返回小扣至少需要多少次操作可以完成所有水缸蓄水要求。

  注意：实际蓄水量 达到或超过 最低蓄水量，即完成蓄水要求。

  示例 1：
    输入：bucket = [1,3], vat = [6,8]
    输出：4
    解释：
    第 1 次操作升级 bucket[0]；
    第 2 ~ 4 次操作均选择蓄水，即可完成蓄水要求。

  示例 2：
    输入：bucket = [9,0,1], vat = [0,2,2]
    输出：3
    解释：
    第 1 次操作均选择升级 bucket[1]
    第 2~3 次操作选择蓄水，即可完成蓄水要求。

  提示：
    1. 1 <= bucket.length == vat.length <= 100
    2. 0 <= bucket[i], vat[i] <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/o8SXZn
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Priority Queue
func storeWater(bucket []int, vat []int) int {
	h := MaxHeap{}
	op := 0
	for i := range bucket {
		if vat[i] == 0 {
			continue
		}
		if bucket[i] == 0 {
			bucket[i] = 1
			op++
		}
		heap.Push(&h, [2]int{bucket[i], vat[i]})
	}

	if h.Len() == 0 {
		return 0
	}

	out := 1<<31 - 1
	for {
		e := heap.Pop(&h).([2]int)
		out = Min(out, op+e[1]/e[0]+remain(e[1], e[0]))
		if e[0] < int(math.Sqrt(float64(e[1]))) {
			op++
			heap.Push(&h, [2]int{e[0] + 1, e[1]})
		} else {
			break
		}
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func remain(a, b int) int {
	if a%b == 0 {
		return 0
	}
	return 1
}

type MaxHeap [][2]int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return float64(h[i][1])/float64(h[i][0]) > float64(h[j][1])/float64(h[j][0])
}
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *MaxHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
