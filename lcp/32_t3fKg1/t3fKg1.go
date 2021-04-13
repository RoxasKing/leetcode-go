package main

import (
	"container/heap"
	"sort"
)

/*
  某实验室计算机待处理任务以 [start,end,period] 格式记于二维数组 tasks，表示完成该任务的时间范围为起始时间 start 至结束时间 end 之间，需要计算机投入 period 的时长，注意：
    1. period 可为不连续时间
    1. 首尾时间均包含在内
  处于开机状态的计算机可同时处理任意多个任务，请返回电脑最少开机多久，可处理完所有任务。

  示例 1：
    输入：tasks = [[1,3,2],[2,5,3],[5,6,2]]
    输出：4
    解释：
      tasks[0] 选择时间点 2、3；
      tasks[1] 选择时间点 2、3、5；
      tasks[2] 选择时间点 5、6；
      因此计算机仅需在时间点 2、3、5、6 四个时刻保持开机即可完成任务。

  示例 2：
    输入：tasks = [[2,3,1],[5,5,1],[5,6,2]]
    输出：3
    解释：
      tasks[0] 选择时间点 2 或 3；
      tasks[1] 选择时间点 5；
      tasks[2] 选择时间点 5、6；
      因此计算机仅需在时间点 2、5、6 或 3、5、6 三个时刻保持开机即可完成任务。

  提示：
    1. 2 <= tasks.length <= 10^5
    2. tasks[i].length == 3
    3. 0 <= tasks[i][0] <= tasks[i][1] <= 10^9
    4. 1 <= tasks[i][2] <= tasks[i][1]-tasks[i][0] + 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/t3fKg1
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Priority Queue(Heap Sort)
func processTasks(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i][0] < tasks[j][0] })
	tasks = append(tasks, []int{1e9 + 1, 1e9 + 1, 1})
	h := MinHeap{}
	out := 0
	for _, t := range tasks {
		start, end, period := t[0], t[1], t[2]
		for h.Len() > 0 && h[0][0]+out < start {
			if h[0][0]+out >= h[0][1] {
				heap.Pop(&h)
			} else {
				out += Min(h[0][1], start) - (h[0][0] + out)
			}
		}
		heap.Push(&h, [2]int{end - period + 1 - out, end + 1})
	}
	return out
}

type MinHeap [][2]int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *MinHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
