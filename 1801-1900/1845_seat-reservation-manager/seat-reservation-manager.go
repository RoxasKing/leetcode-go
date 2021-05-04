package main

import "container/heap"

/*
  Design a system that manages the reservation state of n seats that are numbered from 1 to n.

  Implement the SeatManager class:
    1. SeatManager(int n) Initializes a SeatManager object that will manage n seats numbered from 1 to n. All seats are initially available.
    2. int reserve() Fetches the smallest-numbered unreserved seat, reserves it, and returns its number.
    3. void unreserve(int seatNumber) Unreserves the seat with the given seatNumber.

  Example 1:
    Input
      ["SeatManager", "reserve", "reserve", "unreserve", "reserve", "reserve", "reserve", "reserve", "unreserve"]
      [[5], [], [], [2], [], [], [], [], [5]]
    Output
      [null, 1, 2, null, 2, 3, 4, 5, null]
    Explanation
      SeatManager seatManager = new SeatManager(5); // Initializes a SeatManager with 5 seats.
      seatManager.reserve();    // All seats are available, so return the lowest numbered seat, which is 1.
      seatManager.reserve();    // The available seats are [2,3,4,5], so return the lowest of them, which is 2.
      seatManager.unreserve(2); // Unreserve seat 2, so now the available seats are [2,3,4,5].
      seatManager.reserve();    // The available seats are [2,3,4,5], so return the lowest of them, which is 2.
      seatManager.reserve();    // The available seats are [3,4,5], so return the lowest of them, which is 3.
      seatManager.reserve();    // The available seats are [4,5], so return the lowest of them, which is 4.
      seatManager.reserve();    // The only available seat is seat 5, so return 5.
      seatManager.unreserve(5); // Unreserve seat 5, so now the available seats are [5].

  Constraints:
    1. 1 <= n <= 10^5
    2. 1 <= seatNumber <= n
    3. For each call to reserve, it is guaranteed that there will be at least one unreserved seat.
    4. For each call to unreserve, it is guaranteed that seatNumber will be reserved.
    5. At most 105 calls in total will be made to reserve and unreserve.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/seat-reservation-manager
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Priority Queue(Heap Sort)

type SeatManager struct {
	h      *MinHeap
	inHeap map[int]bool
}

func Constructor(n int) SeatManager {
	h := &MinHeap{}
	inHeap := make(map[int]bool)
	for i := 1; i <= n; i++ {
		heap.Push(h, i)
		inHeap[i] = true
	}
	return SeatManager{
		h:      h,
		inHeap: inHeap,
	}
}

func (this *SeatManager) Reserve() int {
	num := heap.Pop(this.h).(int)
	this.inHeap[num] = false
	return num
}

func (this *SeatManager) Unreserve(seatNumber int) {
	if this.inHeap[seatNumber] {
		return
	}
	this.inHeap[seatNumber] = true
	heap.Push(this.h, seatNumber)
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */

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
