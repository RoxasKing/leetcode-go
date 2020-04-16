package leetcode

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
