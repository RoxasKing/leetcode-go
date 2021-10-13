package main

// Difficulty:
// Hard

// Tags:
// Union-Find

type SummaryRanges struct {
	exist []bool
	uf    UnionFind
	l, r  int
}

func Constructor() SummaryRanges {
	return SummaryRanges{
		exist: make([]bool, 10001),
		uf:    NewUnionFind(10001),
		l:     10000,
		r:     0,
	}
}

func (this *SummaryRanges) AddNum(val int) {
	if this.exist[val] {
		return
	}
	this.exist[val] = true
	if val > 0 && this.exist[val-1] {
		this.uf.Union(val-1, val)
	}
	if val+1 <= 10000 && this.exist[val+1] {
		this.uf.Union(val, val+1)
	}
	if val < this.l {
		this.l = val
	}
	if val > this.r {
		this.r = val
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	out := [][]int{}
	for i := this.l; i <= this.r; i++ {
		if !this.exist[i] {
			continue
		}
		l, r := i, this.uf.FindR(i)
		out = append(out, []int{l, r})
		i = r
	}
	return out
}

type UnionFind struct {
	l, r []int
}

func NewUnionFind(n int) UnionFind {
	l := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = i
		r[i] = i
	}
	return UnionFind{l: l, r: r}
}

func (uf UnionFind) FindL(x int) int {
	if uf.l[x] != x {
		uf.l[x] = uf.FindL(uf.l[x])
	}
	return uf.l[x]
}

func (uf UnionFind) FindR(x int) int {
	if uf.r[x] != x {
		uf.r[x] = uf.FindR(uf.r[x])
	}
	return uf.r[x]
}

func (uf UnionFind) Union(x, y int) {
	xl, xr := uf.FindL(x), uf.FindR(x)
	yl, yr := uf.FindL(y), uf.FindR(y)
	if xr < yl {
		uf.r[xr] = yr
		uf.l[yl] = xl
	} else if yr < xl {
		uf.r[yr] = xr
		uf.l[xl] = yl
	}
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */
