package main

// Difficulty:
// Hard

// Tags:
// Math
// Hash
// Union-Find

func largestComponentSize(nums []int) int {
	primes := []int{}
	isPrime := [1e5 + 1]bool{}
	for i := 2; i <= 1e5; i++ {
		isPrime[i] = true
	}
	for i := 2; i <= 1e5; i++ {
		if isPrime[i] {
			primes = append(primes, i)
			for j := i << 1; j <= 1e5; j += i {
				isPrime[j] = false
			}
		}
	}
	exist := [1e5 + 1]bool{}
	for _, num := range nums {
		exist[num] = true
	}
	uf := NewUnionFind(1e5 + 1)
	for _, prime := range primes {
		pre := -1
		for x := prime; x <= 1e5; x += prime {
			if exist[x] {
				if pre != -1 {
					uf.Union(pre, x)
				}
				pre = x
			}
		}
	}
	freq := [1e5 + 1]int{}
	for _, num := range nums {
		freq[uf.Find(num)]++
	}
	out := 0
	for i := 0; i <= 1e5; i++ {
		if freq[i] > out {
			out = freq[i]
		}
	}
	return out
}

type UnionFind struct {
	parent, size []int
}

func NewUnionFind(n int) UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return UnionFind{parent: parent, size: size}
}

func (uf UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf UnionFind) Union(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}
	if uf.size[x] < uf.size[y] {
		x, y = y, x
	}
	uf.parent[y] = x
	uf.size[x] += uf.size[y]
}
