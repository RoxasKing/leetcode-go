package main

// Tags:
// Priority Queue

type MedianFinder struct {
	max []int
	min []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.max) > len(this.min) {
		if num < this.max[0] {
			this.min = append(this.min, this.max[0])
			up(this.min, smaller)
			this.max[0] = num
			down(this.max, bigger)
		} else {
			this.min = append(this.min, num)
			up(this.min, smaller)
		}
	} else {
		if len(this.min) > 0 && num > this.min[0] {
			this.max = append(this.max, this.min[0])
			up(this.max, bigger)
			this.min[0] = num
			down(this.min, smaller)
		} else {
			this.max = append(this.max, num)
			up(this.max, bigger)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.max) > len(this.min) {
		return float64(this.max[0])
	}
	return float64(this.max[0]+this.min[0]) / 2.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type CompareFunc func(a, b int) bool

func bigger(a, b int) bool {
	return a > b
}

func smaller(a, b int) bool {
	return a < b
}

func up(arr []int, fn CompareFunc) {
	son := len(arr) - 1
	for son > 0 {
		parent := (son - 1) >> 1
		if !fn(arr[son], arr[parent]) {
			return
		}
		arr[son], arr[parent] = arr[parent], arr[son]
		son = parent
	}
}

func down(arr []int, fn CompareFunc) {
	n := len(arr)
	parent := 0
	for parent < n>>1 {
		son := parent<<1 + 1
		if son > n-1 {
			return
		}
		if son+1 < n && fn(arr[son+1], arr[son]) {
			son++
		}
		if !fn(arr[son], arr[parent]) {
			return
		}
		arr[son], arr[parent] = arr[parent], arr[son]
		parent = son
	}
}
