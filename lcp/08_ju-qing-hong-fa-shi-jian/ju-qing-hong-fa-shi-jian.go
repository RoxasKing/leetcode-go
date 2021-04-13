package main

import (
	"container/heap"
	"sort"
)

/*
  在战略游戏中，玩家往往需要发展自己的势力来触发各种新的剧情。一个势力的主要属性有三种，分别是文明等级（C），资源储备（R）以及人口数量（H）。在游戏开始时（第 0 天），三种属性的值均为 0。

  随着游戏进程的进行，每一天玩家的三种属性都会对应增加，我们用一个二维数组 increase 来表示每天的增加情况。这个二维数组的每个元素是一个长度为 3 的一维数组，例如 [[1,2,1],[3,4,2]] 表示第一天三种属性分别增加 1,2,1 而第二天分别增加 3,4,2。

  所有剧情的触发条件也用一个二维数组 requirements 表示。这个二维数组的每个元素是一个长度为 3 的一维数组，对于某个剧情的触发条件 c[i], r[i], h[i]，如果当前 C >= c[i] 且 R >= r[i] 且 H >= h[i] ，则剧情会被触发。

  根据所给信息，请计算每个剧情的触发时间，并以一个数组返回。如果某个剧情不会被触发，则该剧情对应的触发时间为 -1 。

  示例 1：
    输入： increase = [[2,8,4],[2,5,0],[10,9,8]] requirements = [[2,11,3],[15,10,7],[9,17,12],[8,1,14]]
    输出: [2,-1,3,-1]
    解释：
      初始时，C = 0，R = 0，H = 0
      第 1 天，C = 2，R = 8，H = 4
      第 2 天，C = 4，R = 13，H = 4，此时触发剧情 0
      第 3 天，C = 14，R = 22，H = 12，此时触发剧情 2
      剧情 1 和 3 无法触发。

  示例 2：
    输入： increase = [[0,4,5],[4,8,8],[8,6,1],[10,10,0]] requirements = [[12,11,16],[20,2,6],[9,2,6],[10,18,3],[8,14,9]]
    输出: [-1,4,3,3,3]

  示例 3：
    输入： increase = [[1,1,1]] requirements = [[0,0,0]]
    输出: [0]

  限制：
    1. 1 <= increase.length <= 10000
    2. 1 <= requirements.length <= 100000
    3. 0 <= increase[i] <= 10
    4. 0 <= requirements[i] <= 100000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/ju-qing-hong-fa-shi-jian
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func getTriggerTime(increase [][]int, requirements [][]int) []int {
	m := len(increase)
	for i := 1; i < m; i++ {
		increase[i][0] += increase[i-1][0]
		increase[i][1] += increase[i-1][1]
		increase[i][2] += increase[i-1][2]
	}

	out := make([]int, len(requirements))
	for i, req := range requirements {
		if req[0] == 0 && req[1] == 0 && req[2] == 0 {
			continue
		}
		if req[0] > increase[m-1][0] || req[1] > increase[m-1][1] || req[2] > increase[m-1][2] {
			out[i] = -1
			continue
		}
		out[i] = sort.Search(m, func(i int) bool {
			return increase[i][0] >= req[0] && increase[i][1] >= req[1] && increase[i][2] >= req[2]
		}) + 1
	}

	return out
}

// Priority Queue
func getTriggerTime2(increase [][]int, requirements [][]int) []int {
	n := len(requirements)
	pqC := NewPriorityQueue(requirements, 0)
	pqR := NewPriorityQueue(requirements, 1)
	pqH := NewPriorityQueue(requirements, 2)
	c, r, h := 0, 0, 0
	out := make([]int, n)
	for i, req := range requirements {
		if req[0] > c || req[1] > r || req[2] > h {
			heap.Push(pqC, i)
			heap.Push(pqR, i)
			heap.Push(pqH, i)
			out[i] = -1
		}
	}

	for i, inc := range increase {
		c += inc[0]
		r += inc[1]
		h += inc[2]
		for pqC.Len() > 0 && requirements[pqC.idxs[0]][0] <= c {
			idx := heap.Pop(pqC).(int)
			if out[idx] != -1 {
				continue
			}
			if requirements[idx][1] <= r && requirements[idx][2] <= h {
				out[idx] = i + 1
			}
		}
		for pqR.Len() > 0 && requirements[pqR.idxs[0]][1] <= r {
			idx := heap.Pop(pqR).(int)
			if out[idx] != -1 {
				continue
			}
			if requirements[idx][0] <= c && requirements[idx][2] <= h {
				out[idx] = i + 1
			}
		}
		for pqH.Len() > 0 && requirements[pqH.idxs[0]][2] <= h {
			idx := heap.Pop(pqH).(int)
			if out[idx] != -1 {
				continue
			}
			if requirements[idx][0] <= c && requirements[idx][1] <= r {
				out[idx] = i + 1
			}
		}
	}
	return out
}

type PriorityQueue struct {
	reqs [][]int
	idxs []int
	typ  int
}

func NewPriorityQueue(reqs [][]int, typ int) *PriorityQueue {
	return &PriorityQueue{
		reqs: reqs,
		typ:  typ,
	}
}

func (pq PriorityQueue) Len() int { return len(pq.idxs) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq.reqs[pq.idxs[i]][pq.typ] < pq.reqs[pq.idxs[j]][pq.typ]
}
func (pq PriorityQueue) Swap(i, j int) {
	pq.idxs[i], pq.idxs[j] = pq.idxs[j], pq.idxs[i]
}
func (pq *PriorityQueue) Push(x interface{}) { pq.idxs = append(pq.idxs, x.(int)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := pq.Len() - 1
	out := pq.idxs[top]
	pq.idxs = pq.idxs[:top]
	return out
}
