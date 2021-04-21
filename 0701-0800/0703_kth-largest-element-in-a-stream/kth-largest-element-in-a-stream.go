package main

import "container/heap"

/*
  Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.

  Implement KthLargest class:
    KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of integers nums.
    int add(int val) Returns the element representing the kth largest element in the stream.

  Example 1:
  Input
    ["KthLargest", "add", "add", "add", "add", "add"]
    [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
  Output
    [null, 4, 5, 5, 8, 8]
  Explanation
    KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
    kthLargest.add(3);   // return 4
    kthLargest.add(5);   // return 5
    kthLargest.add(10);  // return 5
    kthLargest.add(9);   // return 8
    kthLargest.add(4);   // return 8

  Constraints:
    1 <= k <= 104
    0 <= nums.length <= 104
    -104 <= nums[i] <= 104
    -104 <= val <= 104
    At most 104 calls will be made to add.
    It is guaranteed that there will be at least k elements in the array when you search for the kth element.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/kth-largest-element-in-a-stream
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Priority Queue(Heap Sort)
type KthLargest struct {
	maxh MaxHeap
	minh MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	maxh := MaxHeap{}
	minh := MinHeap(nums)
	heap.Init(&minh)
	for minh.Len() > k-1 {
		heap.Push(&maxh, heap.Pop(&minh))
	}
	return KthLargest{maxh: maxh, minh: minh}
}

func (this *KthLargest) Add(val int) int {
	if this.minh.Len() > 0 && val > this.minh.Top() {
		heap.Push(&this.maxh, heap.Pop(&this.minh))
		heap.Push(&this.minh, val)
	} else {
		heap.Push(&this.maxh, val)
	}
	return this.maxh.Top()
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

type MaxHeap []int

func (mh MaxHeap) Len() int            { return len(mh) }
func (mh MaxHeap) Less(i, j int) bool  { return mh[i] > mh[j] }
func (mh MaxHeap) Swap(i, j int)       { mh[i], mh[j] = mh[j], mh[i] }
func (mh *MaxHeap) Push(x interface{}) { *mh = append(*mh, x.(int)) }
func (mh *MaxHeap) Pop() interface{} {
	last := len(*mh) - 1
	out := (*mh)[last]
	*mh = (*mh)[:last]
	return out
}
func (mh MaxHeap) Top() int { return mh[0] }

type MinHeap []int

func (mh MinHeap) Len() int            { return len(mh) }
func (mh MinHeap) Less(i, j int) bool  { return mh[i] < mh[j] }
func (mh MinHeap) Swap(i, j int)       { mh[i], mh[j] = mh[j], mh[i] }
func (mh *MinHeap) Push(x interface{}) { *mh = append(*mh, x.(int)) }
func (mh *MinHeap) Pop() interface{} {
	last := len(*mh) - 1
	out := (*mh)[last]
	*mh = (*mh)[:last]
	return out
}
func (mh MinHeap) Top() int { return mh[0] }
