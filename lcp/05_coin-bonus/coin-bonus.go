package main

// Tags:
// Segment Tree
func bonus(n int, leadership [][]int, operations [][]int) []int {
	edges := make([][]int, n+1)
	for _, e := range leadership {
		edges[e[0]] = append(edges[e[0]], e[1])
	}

	idxMap := make([][]int, n+1)
	idx := 0
	build(edges, idxMap, 1, &idx)

	sum := make([]int, n*4)
	lazySum := make([]int, n*4)

	out := []int{}
	for _, op := range operations {
		switch op[0] {
		case 1:
			update(sum, lazySum, idxMap[op[1]][0], idxMap[op[1]][0], op[2], 0, n-1, 0)
		case 2:
			update(sum, lazySum, idxMap[op[1]][0], idxMap[op[1]][1], op[2], 0, n-1, 0)
		case 3:
			res := query(sum, lazySum, idxMap[op[1]][0], idxMap[op[1]][1], 0, n-1, 0)
			out = append(out, res)
		}
	}
	return out
}

func build(edges [][]int, idxMap [][]int, member int, idx *int) {
	idxMap[member] = []int{*idx, *idx}
	*idx++
	for _, n := range edges[member] {
		build(edges, idxMap, n, idx)
		idxMap[member][1] = idxMap[n][1]
	}
}

const mod = 1e9 + 7

func update(sum, lazySum []int, opL, opR, opVal int, l, r int, idx int) {
	pushDown(sum, lazySum, l, r, idx)
	if r < opL || opR < l {
		return
	}
	if opL <= l && r <= opR {
		lazySum[idx] = opVal
		pushDown(sum, lazySum, l, r, idx)
		return
	}
	mid := (l + r) >> 1
	update(sum, lazySum, opL, opR, opVal, l, mid, idx<<1+1)
	update(sum, lazySum, opL, opR, opVal, mid+1, r, idx<<1+2)
	sum[idx] = (sum[idx<<1+1] + sum[idx<<1+2]) % mod
}

func query(sum, lazySum []int, opL, opR int, l, r int, idx int) int {
	if r < opL || opR < l {
		return 0
	}
	pushDown(sum, lazySum, l, r, idx)
	if opL <= l && r <= opR {
		return sum[idx]
	}
	mid := (l + r) >> 1
	lSum := query(sum, lazySum, opL, opR, l, mid, idx<<1+1)
	rSum := query(sum, lazySum, opL, opR, mid+1, r, idx<<1+2)
	return (lSum + rSum) % mod
}

func pushDown(sum, lazySum []int, l, r int, idx int) {
	if lazySum[idx] == 0 {
		return
	}
	sum[idx] = (sum[idx] + lazySum[idx]*(r-l+1)) % mod
	if l < r {
		lazySum[idx<<1+1] = (lazySum[idx<<1+1] + lazySum[idx]) % mod
		lazySum[idx<<1+2] = (lazySum[idx<<1+2] + lazySum[idx]) % mod
	}
	lazySum[idx] = 0
}
