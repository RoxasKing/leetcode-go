package main

// Difficulty:
// Medium

// Tags:
// Segment Tree
// Hash

type MyCalendar struct {
	f    map[int][2]bool
	s, t int
}

func Constructor() MyCalendar {
	return MyCalendar{
		f: map[int][2]bool{},
		s: 0,
		t: 1e9 - 1,
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if this.query(start, end-1) {
		return false
	}
	this.update(start, end-1)
	return true
}

func (this *MyCalendar) query(l, r int) bool {
	var dfs func(i, s, t int) bool
	dfs = func(i, s, t int) bool {
		if r < s || t < l {
			return false
		}
		v := this.f[i]
		if l <= s && t <= r {
			if v[1] {
				v[0] = true
			}
			this.f[i] = v
			return v[0]
		}
		if v[1] && s != t {
			lv, rv := this.f[i*2], this.f[i*2+1]
			lv[0], rv[0] = true, true
			lv[1], rv[1] = true, true
			this.f[i*2], this.f[i*2+1] = lv, rv
			v[1] = false
			this.f[i] = v
		}
		m := (s + t) / 2
		return dfs(i*2, s, m) || dfs(i*2+1, m+1, t)
	}
	return dfs(1, this.s, this.t)
}

func (this *MyCalendar) update(l, r int) {
	var dfs func(i, s, t int)
	dfs = func(i, s, t int) {
		if r < s || t < l {
			return
		}
		v := this.f[i]
		if l <= s && t <= r {
			v[0], v[1] = true, true
			this.f[i] = v
			return
		}
		if v[1] && s != t {
			lv, rv := this.f[i*2], this.f[i*2+1]
			lv[0], rv[0] = true, true
			lv[1], rv[1] = true, true
			this.f[i*2], this.f[i*2+1] = lv, rv
			v[1] = false
		}
		m := (s + t) / 2
		dfs(i*2, s, m)
		dfs(i*2+1, m+1, t)
		v[0] = this.f[i*2][0] || this.f[i*2+1][0]
		this.f[i] = v
	}
	dfs(1, this.s, this.t)
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
