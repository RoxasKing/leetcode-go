package main

import "container/heap"

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
