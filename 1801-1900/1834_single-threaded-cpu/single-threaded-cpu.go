package main

import (
	"container/heap"
	"sort"
)

/*
  You are given n​​​​​​ tasks labeled from 0 to n - 1 represented by a 2D integer array tasks, where tasks[i] = [enqueueTimei, processingTimei] means that the i​​​​​​th​​​​ task will be available to process at enqueueTimei and will take processingTimei to finish processing.

  You have a single-threaded CPU that can process at most one task at a time and will act in the following way:
    1. If the CPU is idle and there are no available tasks to process, the CPU remains idle.
    2. If the CPU is idle and there are available tasks, the CPU will choose the one with the shortest processing time. If multiple tasks have the same shortest processing time, it will choose the task with the smallest index.
    3. Once a task is started, the CPU will process the entire task without stopping.
    4. The CPU can finish a task then start a new one instantly.
  Return the order in which the CPU will process the tasks.

  Example 1:
    Input: tasks = [[1,2],[2,4],[3,2],[4,1]]
    Output: [0,2,3,1]
    Explanation: The events go as follows:
      - At time = 1, task 0 is available to process. Available tasks = {0}.
      - Also at time = 1, the idle CPU starts processing task 0. Available tasks = {}.
      - At time = 2, task 1 is available to process. Available tasks = {1}.
      - At time = 3, task 2 is available to process. Available tasks = {1, 2}.
      - Also at time = 3, the CPU finishes task 0 and starts processing task 2 as it is the shortest. Available tasks = {1}.
      - At time = 4, task 3 is available to process. Available tasks = {1, 3}.
      - At time = 5, the CPU finishes task 2 and starts processing task 3 as it is the shortest. Available tasks = {1}.
      - At time = 6, the CPU finishes task 3 and starts processing task 1. Available tasks = {}.
      - At time = 10, the CPU finishes task 1 and becomes idle.

  Example 2:
    Input: tasks = [[7,10],[7,12],[7,5],[7,4],[7,2]]
    Output: [4,3,2,0,1]
    Explanation: The events go as follows:
      - At time = 7, all the tasks become available. Available tasks = {0,1,2,3,4}.
      - Also at time = 7, the idle CPU starts processing task 4. Available tasks = {0,1,2,3}.
      - At time = 9, the CPU finishes task 4 and starts processing task 3. Available tasks = {0,1,2}.
      - At time = 13, the CPU finishes task 3 and starts processing task 2. Available tasks = {0,1}.
      - At time = 18, the CPU finishes task 2 and starts processing task 0. Available tasks = {1}.
      - At time = 28, the CPU finishes task 0 and starts processing task 1. Available tasks = {}.
      - At time = 40, the CPU finishes task 1 and becomes idle.

  Constraints:
    1. tasks.length == n
    2. 1 <= n <= 10^5
    3. 1 <= enqueueTimei, processingTimei <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/single-threaded-cpu
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Priority Queue(Heap Sort)
func getOrder(tasks [][]int) []int {
	n := len(tasks)
	idxs := make([]int, n)
	for i := range idxs {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool {
		a, b := idxs[i], idxs[j]
		if tasks[a][0] != tasks[b][0] {
			return tasks[a][0] < tasks[b][0]
		}
		return a < b
	})
	t := 0
	idx := 0
	out := make([]int, n)
	pq := PriorityQueue{tasks: tasks}
	for len(idxs) > 0 || pq.Len() > 0 {
		for len(idxs) > 0 && tasks[idxs[0]][0] <= t {
			i := idxs[0]
			idxs = idxs[1:]
			heap.Push(&pq, i)
		}
		if pq.Len() == 0 && len(idxs) > 0 {
			t = tasks[idxs[0]][0]
		}
		if pq.Len() > 0 {
			i := heap.Pop(&pq).(int)
			out[idx] = i
			idx++
			t += tasks[i][1]
		}
	}
	return out
}

type PriorityQueue struct {
	tasks [][]int
	idxs  []int
}

func (pq PriorityQueue) Len() int { return len(pq.idxs) }
func (pq PriorityQueue) Less(i, j int) bool {
	a, b := pq.idxs[i], pq.idxs[j]
	if pq.tasks[a][1] != pq.tasks[b][1] {
		return pq.tasks[a][1] < pq.tasks[b][1]
	}
	return a < b
}
func (pq PriorityQueue) Swap(i, j int)       { pq.idxs[i], pq.idxs[j] = pq.idxs[j], pq.idxs[i] }
func (pq *PriorityQueue) Push(x interface{}) { pq.idxs = append(pq.idxs, x.(int)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := pq.Len() - 1
	out := pq.idxs[top]
	pq.idxs = pq.idxs[:top]
	return out
}
