package main

/*
  如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

  例如，
    [2,3,4] 的中位数是 3
    [2,3] 的中位数是 (2 + 3) / 2 = 2.5

  设计一个支持以下两种操作的数据结构：
    1. void addNum(int num) - 从数据流中添加一个整数到数据结构中。
    2. double findMedian() - 返回目前所有元素的中位数。

  示例 1：
    输入：
    ["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
    [[],[1],[2],[],[3],[]]
    输出：[null,null,null,1.50000,null,2.00000]

  示例 2：
    输入：
    ["MedianFinder","addNum","findMedian","addNum","findMedian"]
    [[],[2],[],[3],[]]
    输出：[null,null,2.00000,null,2.50000]

  限制：
    最多会对 addNum、findMedian 进行 50000 次调用。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Priority Queue(Heap Sort)

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
