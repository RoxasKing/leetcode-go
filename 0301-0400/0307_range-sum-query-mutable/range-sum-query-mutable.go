package main

// Difficulty:
// Medium

// Tags:
// Segment Tree

type NumArray struct {
	f    []int
	s, t int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	f := make([]int, n*4)
	build(f, 1, nums, 0, n-1)
	return NumArray{f, 0, n - 1}
}

func (this *NumArray) Update(index int, val int) {
	update(this.f, 1, this.s, this.t, index, index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	return query(this.f, 1, this.s, this.t, left, right)
}

func build(f []int, i int, nums []int, s, t int) {
	if s == t {
		f[i] = nums[s]
		return
	}
	m := (s + t) >> 1
	build(f, i<<1, nums, s, m)
	build(f, i<<1+1, nums, m+1, t)
	f[i] = f[i<<1] + f[i<<1+1]
}

func update(f []int, i, s, t, l, r, v int) {
	if t < l || r < s {
		return
	}
	if l <= s && t <= r {
		f[i] = v
		return
	}
	m := (s + t) >> 1
	update(f, i<<1, s, m, l, r, v)
	update(f, i<<1+1, m+1, t, l, r, v)
	f[i] = f[i<<1] + f[i<<1+1]
}

func query(f []int, i, s, t, l, r int) int {
	if t < l || r < s {
		return 0
	}
	if l <= s && t <= r {
		return f[i]
	}
	m := (s + t) >> 1
	return query(f, i<<1, s, m, l, r) + query(f, i<<1+1, m+1, t, l, r)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
