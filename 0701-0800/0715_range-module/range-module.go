package main

// Difficulty:
// Hard

// Tags:
// Segment Tree

type RangeModule struct {
	f    map[int][2]int8
	s, t int
}

func Constructor() RangeModule {
	return RangeModule{
		f: map[int][2]int8{},
		s: 0,
		t: 1e9 - 2,
	}
}

func (this *RangeModule) AddRange(left int, right int) {
	left, right = left-1, right-2
	var add func(i, s, t int)
	add = func(i, s, t int) {
		if right < s || t < left {
			return
		}
		v := this.f[i]
		if left <= s && t <= right {
			v[0], v[1] = 1, 1
			this.f[i] = v
			return
		}
		if v[1] != 0 && s != t {
			lv, rv := this.f[i*2], this.f[i*2+1]
			lv[0], rv[0] = max(0, v[1]), max(0, v[1])
			lv[1], rv[1] = v[1], v[1]
			this.f[i*2], this.f[i*2+1] = lv, rv
			v[1] = 0
		}
		m := (s + t) / 2
		add(i*2, s, m)
		add(i*2+1, m+1, t)
		v[0] = this.f[i*2][0] & this.f[i*2+1][0]
		this.f[i] = v
	}
	add(1, this.s, this.t)
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	left, right = left-1, right-2
	var query func(i, s, t int) int8
	query = func(i, s, t int) int8 {
		if right < s || t < left {
			return 1
		}
		v := this.f[i]
		if left <= s && t <= right {
			if v[1] != 0 {
				v[0] = max(0, v[1])
			}
			this.f[i] = v
			return v[0]
		}
		if v[1] != 0 && s != t {
			lv, rv := this.f[i*2], this.f[i*2+1]
			lv[0], rv[0] = max(0, v[1]), max(0, v[1])
			lv[1], rv[1] = v[1], v[1]
			this.f[i*2], this.f[i*2+1] = lv, rv
			v[1] = 0
			this.f[i] = v
		}
		m := (s + t) / 2
		return query(i*2, s, m) & query(i*2+1, m+1, t)
	}
	return query(1, this.s, this.t) == 1
}

func (this *RangeModule) RemoveRange(left int, right int) {
	left, right = left-1, right-2
	var del func(i, s, t int)
	del = func(i, s, t int) {
		if right < s || t < left {
			return
		}
		v := this.f[i]
		if left <= s && t <= right {
			v[0], v[1] = 0, -1
			this.f[i] = v
			return
		}
		if v[1] != 0 && s != t {
			lv, rv := this.f[i*2], this.f[i*2+1]
			lv[0], rv[0] = max(0, v[1]), max(0, v[1])
			lv[1], rv[1] = v[1], v[1]
			this.f[i*2], this.f[i*2+1] = lv, rv
			v[1] = 0
		}
		m := (s + t) / 2
		del(i*2, s, m)
		del(i*2+1, m+1, t)
		v[0] = this.f[i*2][0] & this.f[i*2+1][0]
		this.f[i] = v
	}
	del(1, this.s, this.t)
}

func max(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
