package main

/*
  中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。

  例如，
  [2,3,4] 的中位数是 3
  [2,3] 的中位数是 (2 + 3) / 2 = 2.5

  设计一个支持以下两种操作的数据结构：
    void addNum(int num) - 从数据流中添加一个整数到数据结构中。
    double findMedian() - 返回目前所有元素的中位数。

  进阶:
    如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
    如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-median-from-data-stream
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type MedianFinder struct {
	max []int // 最大堆，其根节点小于等于最小堆的根节点
	min []int // 最小堆，其根节点大于等于最大堆的根节点
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		max: []int{},
		min: []int{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.max) > len(this.min) {
		if num < this.max[0] {
			this.min = append(this.min, this.max[0])
			this.max[0] = num
			down(this.max, bigger)
		} else {
			this.min = append(this.min, num)
		}
		up(this.min, smaller)
	} else {
		if len(this.min) != 0 && num > this.min[0] {
			this.max = append(this.max, this.min[0])
			this.min[0] = num
			down(this.min, smaller)
		} else {
			this.max = append(this.max, num)
		}
		up(this.max, bigger)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.max) == len(this.min) {
		return float64(this.max[0]+this.min[0]) / 2
	}
	return float64(this.max[0])
}

// 节点升顶
func up(nums []int, compare compareFunc) {
	son := len(nums) - 1
	for son > 0 {
		parent := (son - 1) >> 1
		if !compare(nums, son, parent) {
			return
		}
		nums[son], nums[parent] = nums[parent], nums[son]
		son = parent
	}
}

// 节点沉降
func down(nums []int, compare compareFunc) {
	n := len(nums)
	var parent int
	for parent < n>>1 {
		son := parent<<1 + 1
		if son > n-1 {
			return
		}
		if son+1 < n && compare(nums, son+1, son) {
			son++
		}
		if !compare(nums, son, parent) {
			return
		}
		nums[son], nums[parent] = nums[parent], nums[son]
		parent = son
	}
}

type compareFunc func([]int, int, int) bool

func smaller(nums []int, i, j int) bool {
	return nums[i] < nums[j]
}

func bigger(nums []int, i, j int) bool {
	return nums[i] > nums[j]
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
