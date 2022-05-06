package main

// Difficulty:
// Easy

// Tags:
// Queue

type MyStack struct {
	a, b []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.a = append(this.a, x)
}

func (this *MyStack) Pop() int {
	if len(this.a) == 0 {
		for ; len(this.b) > 0; this.b = this.b[1:] {
			this.a = append(this.a, this.b[0])
		}
	}
	for ; len(this.a) > 1; this.a = this.a[1:] {
		this.b = append(this.b, this.a[0])
	}
	o := this.a[0]
	this.a = this.a[1:]
	return o
}

func (this *MyStack) Top() int {
	if len(this.a) == 0 {
		for ; len(this.b) > 0; this.b = this.b[1:] {
			this.a = append(this.a, this.b[0])
		}
	}
	for ; len(this.a) > 1; this.a = this.a[1:] {
		this.b = append(this.b, this.a[0])
	}
	return this.a[0]
}

func (this *MyStack) Empty() bool {
	return len(this.a) == 0 && len(this.b) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
