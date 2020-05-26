package codinginterviews

/*
  如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

  例如，
    [2,3,4] 的中位数是 3
    [2,3] 的中位数是 (2 + 3) / 2 = 2.5

  设计一个支持以下两种操作的数据结构：
    void addNum(int num) - 从数据流中添加一个整数到数据结构中。
    double findMedian() - 返回目前所有元素的中位数。

  限制：
    最多会对 addNum、findMedia进行 50000 次调用。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type MedianFinder struct {
	maxHeap []int
	minHeap []int
}

func maxHeapAjust(heap []int) {
	for i := len(heap)>>1 - 1; i >= 0; i-- {
		son := i<<1 + 1
		if son > len(heap)-1 {
			continue
		}
		if son < len(heap)-1 && heap[son+1] > heap[son] {
			son++
		}
		if heap[son] > heap[i] {
			heap[son], heap[i] = heap[i], heap[son]
		}
	}
}

func minHeapAjust(heap []int) {
	for i := len(heap)>>1 - 1; i >= 0; i-- {
		son := i<<1 + 1
		if son > len(heap)-1 {
			continue
		}
		if son < len(heap)-1 && heap[son+1] < heap[son] {
			son++
		}
		if heap[son] < heap[i] {
			heap[son], heap[i] = heap[i], heap[son]
		}
	}
}

/** initialize your data structure here. */
func NewMedianFinder() MedianFinder {
	return MedianFinder{
		maxHeap: []int{},
		minHeap: []int{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.maxHeap) > len(this.minHeap) {
		if num < this.maxHeap[0] {
			num, this.maxHeap[0] = this.maxHeap[0], num
			maxHeapAjust(this.maxHeap)
		}
		this.minHeap = append(this.minHeap, num)
		minHeapAjust(this.minHeap)
	} else {
		if len(this.minHeap) > 0 && num > this.minHeap[0] {
			num, this.minHeap[0] = this.minHeap[0], num
			minHeapAjust(this.minHeap)
		}
		this.maxHeap = append(this.maxHeap, num)
		maxHeapAjust(this.maxHeap)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.maxHeap) > len(this.minHeap) {
		return float64(this.maxHeap[0])
	}
	return float64(this.maxHeap[0]+this.minHeap[0]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
