package main

import (
	"container/heap"
	"math/rand"
)

/*
  给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

  提示：
    你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
    你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
    题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
    你可以按任意顺序返回答案。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/top-k-frequent-elements
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Heap Sort + Min Heap + Use Container
func topKFrequent(nums []int, k int) []int {
	rand.Intn(len(nums))
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	q := PriorityQueue{}
	for k, v := range count {
		heap.Push(&q, &elem{key: k, count: v})
	}
	out := make([]int, 0, k)
	for len(out) < k {
		e := heap.Pop(&q).(*elem)
		out = append(out, e.key)
	}
	return out
}

type elem struct {
	key   int
	count int
}

type PriorityQueue []*elem

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].count > p[j].count
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p *PriorityQueue) Push(x interface{}) {
	*p = append(*p, x.(*elem))
}

func (p *PriorityQueue) Pop() interface{} {
	e := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return e
}

// Heap Sort + Min Heap
func topKFrequent2(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	var heap []int
	for key := range count {
		heap = append(heap, key)
		up(heap, count)
	}
	out := make([]int, 0, k)
	for len(out) < k {
		last := len(heap) - 1
		heap[0], heap[last] = heap[last], heap[0]
		out = append(out, heap[last])
		heap = heap[:last]
		down(heap, count)
	}
	return out
}

func up(nums []int, count map[int]int) {
	son := len(nums) - 1
	for son > 0 {
		parent := (son - 1) >> 1
		if count[nums[parent]] >= count[nums[son]] {
			return
		}
		nums[parent], nums[son] = nums[son], nums[parent]
		son = parent
	}
}

func down(nums []int, count map[int]int) {
	var parent int
	for parent < len(nums)>>1 {
		son := parent*2 + 1
		if son > len(nums)-1 {
			return
		}
		if son+1 < len(nums) && count[nums[son+1]] > count[nums[son]] {
			son++
		}
		if count[nums[son]] <= count[nums[parent]] {
			return
		}
		nums[parent], nums[son] = nums[son], nums[parent]
		parent = son
	}
}
