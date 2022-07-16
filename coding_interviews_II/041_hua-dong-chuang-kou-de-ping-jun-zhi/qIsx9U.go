package main

// Difficulty:
// Easy

// Tags:
// Sliding Window

type MovingAverage struct {
	a    []int
	x, c int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{[]int{}, 0, size}
}

func (this *MovingAverage) Next(val int) float64 {
	this.a = append(this.a, val)
	this.x += val
	if len(this.a) > this.c {
		this.x -= this.a[0]
		this.a = this.a[1:]
	}
	return float64(this.x) / float64(len(this.a))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
