package main

// Difficulty:
// Hard

// Tags:
// Segment Tree

type val struct{ a, b int8 }
type tree map[int]val

type MyCalendarThree struct {
	t tree
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{t: map[int]val{}}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	var update func(i, s, t int)
	update = func(i, s, t int) {
		if t < start || end < s {
			return
		}
		p := this.t[i]
		if start <= s && t <= end {
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
	end-- // Attention!
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
