package main

/*
  有 N 个网络节点，标记为 1 到 N。
  给定一个列表 times，表示信号经过有向边的传递时间。 times[i] = (u, v, w)，其中 u 是源节点，v 是目标节点， w 是一个信号从源节点传递到目标节点的时间。
  现在，我们从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1。
*/

// DFS
func networkDelayTime(times [][]int, N int, K int) int {
	graph := make(map[int][][]int)
	for _, t := range times {
		graph[t[0]] = append(graph[t[0]], []int{t[1], t[2]})
	}
	dist := make(map[int]int)
	for i := 1; i <= N; i++ {
		dist[i] = 1<<31 - 1
	}
	dfsNetworkDelayTime(&graph, &dist, K, 0)
	var out int
	for _, v := range dist {
		if v == 1<<31-1 {
			return -1
		}
		out = Max(out, v)
	}
	return out
}

func dfsNetworkDelayTime(graph *map[int][][]int, dist *map[int]int, node int, cur int) {
	if cur >= (*dist)[node] {
		return
	}
	(*dist)[node] = cur
	if _, ok := (*graph)[node]; ok {
		for _, elem := range (*graph)[node] {
			dfsNetworkDelayTime(graph, dist, elem[0], cur+elem[1])
		}
	}
}

// Dijkstra's + Heap
func networkDelayTime2(times [][]int, N int, K int) int {
	graph := make(map[int][][]int)
	for _, t := range times {
		graph[t[0]] = append(graph[t[0]], []int{t[1], t[2]})
	}
	dist := make(map[int]int)
	heap := [][]int{{K, 0}}
	ajust := func() {
		for i := len(heap)/2 - 1; i >= 0; i-- {
			son := 2*i + 1
			if son > len(heap)-1 {
				return
			}
			if son+1 < len(heap) && heap[son+1][1] < heap[son][1] {
				son++
			}
			if heap[son][1] < heap[i][1] {
				heap[i], heap[son] = heap[son], heap[i]
			}
		}
	}
	for len(heap) != 0 {
		ajust()
		node, cur := heap[0][0], heap[0][1]
		heap = heap[1:]
		if _, ok := dist[node]; ok {
			continue
		}
		dist[node] = cur
		if _, ok := graph[node]; !ok {
			continue
		}
		for _, edge := range graph[node] {
			if _, ok := dist[edge[0]]; ok {
				continue
			}
			heap = append(heap, []int{edge[0], cur + edge[1]})
		}
	}
	if len(dist) != N {
		return -1
	}
	var out int
	for _, v := range dist {
		out = Max(out, v)
	}
	return out
}

func networkDelayTime3(times [][]int, N int, K int) int {
	graph := make([][]int, N)
	for i := range graph {
		graph[i] = make([]int, N)
		for j := range graph[0] {
			graph[i][j] = -1
		}
	}
	for _, t := range times {
		graph[t[0]-1][t[1]-1] = t[2]
	}
	mark := make([]bool, N)
	dist := make([]int, N)
	for i := range dist {
		dist[i] = 1<<31 - 1
	}
	dist[K-1] = 0
	minDist := func() int {
		var node int
		min := 1<<31 - 1
		for i := range dist {
			if !mark[i] && dist[i] < min {
				min = dist[i]
				node = i
			}
		}
		return node
	}
	for n := 0; n < N-1; n++ {
		i := minDist()
		mark[i] = true
		if dist[i] == 1<<31-1 {
			continue
		}
		for j := 0; j < N; j++ {
			if mark[j] || graph[i][j] == -1 {
				continue
			}
			dist[j] = Min(dist[j], dist[i]+graph[i][j])
		}
	}
	var out int
	for _, d := range dist {
		out = Max(out, d)
	}
	if out == 1<<31-1 {
		return -1
	}
	return out
}
