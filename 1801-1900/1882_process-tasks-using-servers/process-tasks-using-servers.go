package main

import "container/heap"

/*
  You are given two 0-indexed integer arrays servers and tasks of lengths n​​​​​​ and m​​​​​​ respectively. servers[i] is the weight of the i​​​​​​th​​​​ server, and tasks[j] is the time needed to process the j​​​​​​th​​​​ task in seconds.

  You are running a simulation system that will shut down after all tasks are processed. Each server can only process one task at a time. You will be able to process the jth task starting from the jth second beginning with the 0th task at second 0. To process task j, you assign it to the server with the smallest weight that is free, and in case of a tie, choose the server with the smallest index. If a free server gets assigned task j at second t,​​​​​​ it will be free again at the second t + tasks[j].

  If there are no free servers, you must wait until one is free and execute the free tasks as soon as possible. If multiple tasks need to be assigned, assign them in order of increasing index.

  You may assign multiple tasks at the same second if there are multiple free servers.

  Build an array ans​​​​ of length m, where ans[j] is the index of the server the j​​​​​​th task will be assigned to.

  Return the array ans​​​​.

  Example 1:
    Input: servers = [3,3,2], tasks = [1,2,3,2,1,2]
    Output: [2,2,0,2,1,2]
    Explanation: Events in chronological order go as follows:
      - At second 0, task 0 is added and processed using server 2 until second 1.
      - At second 1, server 2 becomes free. Task 1 is added and processed using server 2 until second 3.
      - At second 2, task 2 is added and processed using server 0 until second 5.
      - At second 3, server 2 becomes free. Task 3 is added and processed using server 2 until second 5.
      - At second 4, task 4 is added and processed using server 1 until second 5.
      - At second 5, all servers become free. Task 5 is added and processed using server 2 until second 7.

  Example 2:
    Input: servers = [5,1,4,3,2], tasks = [2,1,2,4,5,2,1]
    Output: [1,4,1,4,1,3,2]
    Explanation: Events in chronological order go as follows:
      - At second 0, task 0 is added and processed using server 1 until second 2.
      - At second 1, task 1 is added and processed using server 4 until second 2.
      - At second 2, servers 1 and 4 become free. Task 2 is added and processed using server 1 until second 4.
      - At second 3, task 3 is added and processed using server 4 until second 7.
      - At second 4, server 1 becomes free. Task 4 is added and processed using server 1 until second 9.
      - At second 5, task 5 is added and processed using server 3 until second 7.
      - At second 6, task 6 is added and processed using server 2 until second 7.

  Constraints:
    1. servers.length == n
    2. tasks.length == m
    3. 1 <= n, m <= 2 * 10^5
    4. 1 <= servers[i], tasks[j] <= 2 * 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/process-tasks-using-servers
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Priority Queue

func assignTasks(servers []int, tasks []int) []int {
	n := len(tasks)
	tq := [][2]int{}
	for i, t := range tasks {
		tq = append(tq, [2]int{i, t})
	}
	fq := freeQueue{}
	for i, w := range servers {
		heap.Push(&fq, &freeServer{id: i, weight: w})
	}
	wq := workQueue{}
	out := make([]int, n)
	for t := 0; len(tq) > 0; t++ {
		for wq.Len() > 0 && wq[0].stop <= t {
			id := heap.Pop(&wq).(*workServer).id
			heap.Push(&fq, &freeServer{id: id, weight: servers[id]})
		}

		if fq.Len() == 0 {
			t = wq[0].stop - 1
			continue
		}

		for fq.Len() > 0 && len(tq) > 0 && tq[0][0] <= t {
			id := heap.Pop(&fq).(*freeServer).id
			out[tq[0][0]] = id
			heap.Push(&wq, &workServer{id: id, stop: t + tq[0][1]})
			tq = tq[1:]
		}
	}
	return out
}

type freeServer struct {
	id, weight int
}

type freeQueue []*freeServer

func (s freeQueue) Len() int { return len(s) }
func (s freeQueue) Less(i, j int) bool {
	if s[i].weight != s[j].weight {
		return s[i].weight < s[j].weight
	}
	return s[i].id < s[j].id
}
func (s freeQueue) Swap(i, j int)       { s[i], s[j] = s[j], s[i] }
func (s *freeQueue) Push(x interface{}) { *s = append(*s, x.(*freeServer)) }
func (s *freeQueue) Pop() interface{} {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}

type workServer struct {
	id, stop int
}

type workQueue []*workServer

func (s workQueue) Len() int            { return len(s) }
func (s workQueue) Less(i, j int) bool  { return s[i].stop < s[j].stop }
func (s workQueue) Swap(i, j int)       { s[i], s[j] = s[j], s[i] }
func (s *workQueue) Push(x interface{}) { *s = append(*s, x.(*workServer)) }
func (s *workQueue) Pop() interface{} {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}
