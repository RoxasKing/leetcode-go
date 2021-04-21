package main

/*
  There is an undirected graph with n nodes, where each node is numbered between 0 and n - 1. You are given a 2D array graph, where graph[u] is an array of nodes that node u is adjacent to. More formally, for each v in graph[u], there is an undirected edge between node u and node v. The graph has the following properties:
    1. There are no self-edges (graph[u] does not contain u).
    2. There are no parallel edges (graph[u] does not contain duplicate values).
    3. If v is in graph[u], then u is in graph[v] (the graph is undirected).
    4. The graph may not be connected, meaning there may be two nodes u and v such that there is no path between them.
  A graph is bipartite if the nodes can be partitioned into two independent sets A and B such that every edge in the graph connects a node in set A and a node in set B.

  Return true if and only if it is bipartite.

  Example 1:
    Input: graph = [[1,2,3],[0,2],[0,1,3],[0,2]]
    Output: false
    Explanation: There is no way to partition the nodes into two independent sets such that every edge connects a node in one and a node in the other.

  Example 2:
    Input: graph = [[1,3],[0,2],[1,3],[0,2]]
    Output: true
    Explanation: We can partition the nodes into two sets: {0, 2} and {1, 3}.

  Constraints:
    1. graph.length == n
    2. 1 <= n <= 100
    3. 0 <= graph[u].length < n
    4. 0 <= graph[u][i] <= n - 1
    5. graph[u] does not contain u.
    6. All the values of graph[u] are unique.
    7. If graph[u] contains v, then graph[v] contains u.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/is-graph-bipartite
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Bipartite Graph + DFS
func isBipartite(graph [][]int) bool {
	n := len(graph)
	colors := make([]int, n)
	for i := 0; i < n; i++ {
		if colors[i] != 0 {
			continue
		}
		if ok := dfs(graph, colors, i, 1); !ok {
			return false
		}
	}
	return true
}

func dfs(graph [][]int, colors []int, node, color int) bool {
	colors[node] = color
	nextColor := 1
	if color == 1 {
		nextColor = 2
	}
	for _, neighbor := range graph[node] {
		if colors[neighbor] == 0 {
			if ok := dfs(graph, colors, neighbor, nextColor); !ok {
				return false
			}
		} else if colors[neighbor] != nextColor {
			return false
		}
	}
	return true
}

// Bipartite Graph + BFS
func isBipartite2(graph [][]int) bool {
	const UNCOLORED, RED, GREEN = 0, 1, 2
	colors := make([]int, len(graph))
	for i := range graph {
		if colors[i] == UNCOLORED {
			queue := []int{i}
			colors[i] = RED
			for len(queue) != 0 {
				node := queue[0]
				queue = queue[1:]
				nextColor := RED
				if colors[node] == RED {
					nextColor = GREEN
				}
				for _, neighbor := range graph[node] {
					if colors[neighbor] == UNCOLORED {
						queue = append(queue, neighbor)
						colors[neighbor] = nextColor
					} else if colors[neighbor] != nextColor {
						return false
					}
				}
			}
		}
	}
	return true
}
