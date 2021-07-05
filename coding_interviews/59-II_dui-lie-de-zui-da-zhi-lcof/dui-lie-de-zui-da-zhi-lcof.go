package main

// Tags:
// Monotonous Queue

type MaxQueue struct {
	q []int // queue
	m []int // monotonous queue
}

func Constructor() MaxQueue {
	return MaxQueue{
		q: []int{},
		m: []int{},
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.m) == 0 {
		return -1
	}
	return this.m[0]
}

func (this *MaxQueue) Push_back(value int) {
	for len(this.m) > 0 && this.m[len(this.m)-1] < value {
		this.m = this.m[:len(this.m)-1]
	}
	this.m = append(this.m, value)
	this.q = append(this.q, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.m) == 0 {
		return -1
	}
	out := this.q[0]
	this.q = this.q[1:]
	if out == this.m[0] {
		this.m = this.m[1:]
	}
	return out
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
