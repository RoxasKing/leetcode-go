package main

import "container/heap"

/*
  给你一个按升序排序的整数数组 num（可能包含重复数字），请你将它们分割成一个或多个子序列，其中每个子序列都由连续整数组成且长度至少为 3 。
  如果可以完成上述分割，则返回 true ；否则，返回 false 。

  示例 1：
    输入: [1,2,3,3,4,5]
    输出: True
    解释:
    你可以分割出这样两个连续子序列 :
    1, 2, 3
    3, 4, 5

  示例 2：
    输入: [1,2,3,3,4,4,5,5]
    输出: True
    解释:
    你可以分割出这样两个连续子序列 :
    1, 2, 3, 4, 5
    3, 4, 5

  示例 3：
    输入: [1,2,3,4,4,5]
    输出: False

  提示：
    输入的数组长度范围为 [1, 10000]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Priority Queue + Hash
func isPossible(nums []int) bool {
	pqs := make(map[int]*priorityQueue)
	for _, num := range nums {
		if pqs[num] == nil {
			pqs[num] = new(priorityQueue)
		}
		if pq := pqs[num-1]; pq != nil {
			prevLen := heap.Pop(pq).(int)
			if pq.Len() == 0 {
				delete(pqs, num-1)
			}
			heap.Push(pqs[num], prevLen+1)
		} else {
			heap.Push(pqs[num], 1)
		}
	}

	for _, pq := range pqs {
		if (*pq)[0] < 3 {
			return false
		}
	}

	return true
}

type priorityQueue []int

func (p priorityQueue) Len() int            { return len(p) }
func (p priorityQueue) Less(i, j int) bool  { return p[i] < p[j] }
func (p priorityQueue) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *priorityQueue) Push(x interface{}) { *p = append(*p, x.(int)) }
func (p *priorityQueue) Pop() interface{} {
	last := len(*p) - 1
	out := (*p)[last]
	*p = (*p)[:last]
	return out
}
