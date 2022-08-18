package main

// Difficulty:
// Easy

// Tags:
// Two Pointers

type OrderedStream struct {
	a    []string
	l, r int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{make([]string, n), 0, 0}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.a[idKey-1] = value
	for ; this.r < len(this.a) && this.a[this.r] != ""; this.r++ {
	}
	defer func() { this.l = this.r }()
	return this.a[this.l:this.r]
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
