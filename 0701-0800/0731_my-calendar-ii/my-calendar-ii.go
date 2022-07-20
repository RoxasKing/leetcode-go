package main

// Difficulty:
// Medium

// Tags:
// Segment Tree

type MyCalendarTwo struct {
	f    map[int][2]int8
	s, t int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{map[int][2]int8{}, 0, 1e9 - 1}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	l, r := start, end-1
	var add func(i, s, t int, x int8)
	add = func(i, s, t int, x int8) {
		if r < s || t < l {
			return
		}
		v := this.f[i]
		if l <= s && t <= r {
			v[0] += x
			v[1] += x
			this.f[i] = v
			return
		}
		if v[1] != 0 && s != t {
			lv, rv := this.f[i*2], this.f[i*2+1]
			lv[0] += v[1]
			lv[1] += v[1]
			rv[0] += v[1]
			rv[1] += v[1]
			this.f[i*2], this.f[i*2+1] = lv, rv
			v[1] = 0
			this.f[i] = v
		}
		m := (s + t) / 2
		add(i*2, s, m, x)
		add(i*2+1, m+1, t, x)
		v[0] = max(this.f[i*2][0], this.f[i*2+1][0])
		this.f[i] = v
	}
	if add(1, this.s, this.t, 1); this.f[1][0] > 2 {
		add(1, this.s, this.t, -1)
		return false
	}
	return true
}

func max(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
