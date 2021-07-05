package main

// Tags:
// Topological sorting
// Math

func waysToBuildRooms(prevRoom []int) int {
	n := len(prevRoom)
	fac := make([]int, n+1)
	ifac := make([]int, n+1)
	fac[0], fac[1] = 1, 1
	ifac[0], ifac[1] = 1, 1
	for i := 2; i <= n; i++ {
		fac[i] = fac[i-1] * i % mod
		ifac[i] = quickMul(fac[i], mod-2)
	}

	edges := make([][]int, n)
	indeg := make([]int, n)
	for i := 1; i < n; i++ {
		edges[prevRoom[i]] = append(edges[prevRoom[i]], i)
		indeg[prevRoom[i]]++
	}

	lens := make([]int, n)
	for i := range lens {
		lens[i] = 1
	}

	q := []int{}
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	f := make([]int, n)

	for ; len(q) > 0; q = q[1:] {
		u := q[0]
		p := prevRoom[u]
		if 0 <= p {
			lens[p] += lens[u]
			indeg[p]--
			if indeg[p] == 0 {
				q = append(q, p)
			}
		}
		cnt := fac[lens[u]-1]
		for _, v := range edges[u] {
			cnt = cnt * ifac[lens[v]] % mod
			cnt = cnt * f[v] % mod
		}
		f[u] = cnt
	}

	return f[0]
}

func quickMul(x, n int) int {
	out := 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			out = out * x % mod
		}
		x = x * x % mod
	}
	return out
}

var mod = int(1e9 + 7)
