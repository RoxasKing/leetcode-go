package main

// Tags:
// BFS
// Math

func getCoprimes(nums []int, edges [][]int) []int {
	n := len(nums)
	es := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		es[u] = append(es[u], v)
		es[v] = append(es[v], u)
	}

	out := make([]int, n)
	dep := make([]int, n)
	for i := range out {
		out[i] = -1
		dep[i] = -1
	}

	visited := make([]bool, n)
	parents := make([][51][2]int, n)
	for i := range parents {
		parents[i] = [51][2]int{}
		for j := 1; j <= 50; j++ {
			parents[i][j] = [2]int{-1, -1}
		}
	}

	visited[0] = true
	q := [][2]int{{0, 0}}
	for len(q) > 0 {
		u, depth := q[0][0], q[0][1]
		q = q[1:]
		for i := 1; i <= 50; i++ {
			if parents[u][i][0] != -1 && Gcd(nums[u], i) == 1 && (parents[u][i][1] > dep[u]) {
				out[u] = parents[u][i][0]
				dep[u] = parents[u][i][1]
			}
		}
		parents[u][nums[u]] = [2]int{u, depth}
		for _, v := range es[u] {
			if visited[v] {
				continue
			}
			visited[v] = true
			parents[v] = parents[u]
			q = append(q, [2]int{v, depth + 1})
		}
	}

	return out
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
