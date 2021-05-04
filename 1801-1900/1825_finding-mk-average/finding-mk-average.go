package main

import "container/heap"

/*
  You are given two integers, m and k, and a stream of integers. You are tasked to implement a data structure that calculates the MKAverage for the stream.

  The MKAverage can be calculated using these steps:
    1. If the number of the elements in the stream is less than m you should consider the MKAverage to be -1. Otherwise, copy the last m elements of the stream to a separate container.
    2. Remove the smallest k elements and the largest k elements from the container.
    3. Calculate the average value for the rest of the elements rounded down to the nearest integer.
  Implement the MKAverage class:
    1. MKAverage(int m, int k) Initializes the MKAverage object with an empty stream and the two integers m and k.
    2. void addElement(int num) Inserts a new element num into the stream.
    3. int calculateMKAverage() Calculates and returns the MKAverage for the current stream rounded down to the nearest integer.

  Example 1:
    Input
      ["MKAverage", "addElement", "addElement", "calculateMKAverage", "addElement", "calculateMKAverage", "addElement", "addElement", "addElement", "calculateMKAverage"]
      [[3, 1], [3], [1], [], [10], [], [5], [5], [5], []]
    Output
      [null, null, null, -1, null, 3, null, null, null, 5]
    Explanation
      MKAverage obj = new MKAverage(3, 1);
      obj.addElement(3);        // current elements are [3]
      obj.addElement(1);        // current elements are [3,1]
      obj.calculateMKAverage(); // return -1, because m = 3 and only 2 elements exist.
      obj.addElement(10);       // current elements are [3,1,10]
      obj.calculateMKAverage(); // The last 3 elements are [3,1,10].
                                // After removing smallest and largest 1 element the container will be [3].
                                // The average of [3] equals 3/1 = 3, return 3
      obj.addElement(5);        // current elements are [3,1,10,5]
      obj.addElement(5);        // current elements are [3,1,10,5,5]
      obj.addElement(5);        // current elements are [3,1,10,5,5,5]
      obj.calculateMKAverage(); // The last 3 elements are [5,5,5].
                                // After removing smallest and largest 1 element the container will be [5].
                                // The average of [5] equals 5/1 = 5, return 5

  Constraints:
    1. 3 <= m <= 10^5
    2. 1 <= k*2 < m
    3. 1 <= num <= 10^5
    4. At most 10^5 calls will be made to addElement and calculateMKAverage.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/finding-mk-average
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Priority Queue(Heap Sort)
type MKAverage struct {
	m, k int
	arr  []int       // record last m number
	sum  int         // sum of last m number
	maxA MaxHeap     // min k's number
	minA MinHeap     // max (m-k)'s number
	minS int         // min k's sum
	delA map[int]int // nums in min heap needs to be deleted
	minD int         // count of nums in min heap needs to be removed
	minB MinHeap     // max k's number
	maxB MaxHeap     // min (m-k)'s number
	maxS int         // max k's sum
	delB map[int]int // nums in max heap needs to be deleted
	maxD int         // count of nums in max heap needs to be removed
}

func Constructor(m int, k int) MKAverage {
	return MKAverage{
		m:    m,
		k:    k,
		arr:  []int{},
		sum:  0,
		minA: MinHeap{},
		maxA: MaxHeap{},
		maxS: 0,
		delA: make(map[int]int),
		minD: 0,
		maxB: MaxHeap{},
		minB: MinHeap{},
		minS: 0,
		delB: make(map[int]int),
		maxD: 0,
	}
}

func (this *MKAverage) AddElement(num int) {
	if len(this.arr) == this.m {
		num := this.arr[0]
		this.arr = this.arr[1:]
		this.sum -= num
		this.delA[num]++
		if num <= this.maxA[0] {
			this.minS -= num
			this.minD++
		}
		this.delB[num]++
		if num >= this.minB[0] {
			this.maxS -= num
			this.maxD++
		}
	}
	this.arr = append(this.arr, num)
	this.sum += num
	if this.maxA.Len() == 0 || num <= this.maxA[0] {
		heap.Push(&this.maxA, num)
		this.minS += num
	} else {
		heap.Push(&this.minA, num)
	}
	if this.minB.Len() == 0 || num >= this.minB[0] {
		heap.Push(&this.minB, num)
		this.maxS += num
	} else {
		heap.Push(&this.maxB, num)
	}
	this.balanceHeapA()
	this.balanceHeapB()
}

func (this *MKAverage) CalculateMKAverage() int {
	if len(this.arr) < this.m {
		return -1
	}
	return (this.sum - this.minS - this.maxS) / (this.m - this.k*2)
}

/**
 * Your MKAverage object will be instantiated and called as such:
 * obj := Constructor(m, k);
 * obj.AddElement(num);
 * param_2 := obj.CalculateMKAverage();
 */

func (this *MKAverage) balanceHeapA() {
	for {
		for this.maxA.Len() > 0 && this.delA[this.maxA[0]] > 0 {
			this.delA[heap.Pop(&this.maxA).(int)]--
			this.minD--
		}
		for this.minA.Len() > 0 && this.delA[this.minA[0]] > 0 {
			this.delA[heap.Pop(&this.minA).(int)]--
		}
		if this.maxA.Len()-this.minD < this.k && this.minA.Len() > 0 {
			num := heap.Pop(&this.minA).(int)
			this.minS += num
			heap.Push(&this.maxA, num)
		} else if this.maxA.Len()-this.minD > this.k {
			num := heap.Pop(&this.maxA).(int)
			this.minS -= num
			heap.Push(&this.minA, num)
		} else {
			break
		}
	}
}

func (this *MKAverage) balanceHeapB() {
	for {
		for this.minB.Len() > 0 && this.delB[this.minB[0]] > 0 {
			this.delB[heap.Pop(&this.minB).(int)]--
			this.maxD--
		}
		for this.maxB.Len() > 0 && this.delB[this.maxB[0]] > 0 {
			this.delB[heap.Pop(&this.maxB).(int)]--
		}
		if this.minB.Len()-this.maxD < this.k && this.maxB.Len() > 0 {
			num := heap.Pop(&this.maxB).(int)
			this.maxS += num
			heap.Push(&this.minB, num)
		} else if this.minB.Len()-this.maxD > this.k {
			num := heap.Pop(&this.minB).(int)
			this.maxS -= num
			heap.Push(&this.maxB, num)
		} else {
			break
		}
	}
}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	last := len(*h) - 1
	out := (*h)[last]
	*h = (*h)[:last]
	return out
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	last := len(*h) - 1
	out := (*h)[last]
	*h = (*h)[:last]
	return out
}
