package main

// Difficulty:
// Hard

// Tags:
// Segment Tree

type pair struct{ a, b int8 }

type MyCalendarThree struct {
	t map[int]pair
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{map[int]pair{}}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	l, r := start, end-1
	var update func(i, s, t int)
	update = func(i, s, t int) {
		if r < s || t < l {
			return
		}
		p := this.t[i]
		if l <= s && t <= r {
			p.a++
			p.b++
			this.t[i] = p
			return
		}
		m := (s + t) / 2
		update(i*2, s, m)
		update(i*2+1, m+1, t)
		p.a = p.b + max(this.t[i*2].a, this.t[i*2+1].a)
		this.t[i] = p
	}
	update(1, 0, 1e9)
	return int(this.t[1].a)
}

func max(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
