package main

import "container/heap"

/*
  Design a stack-like data structure to push elements to the stack and pop the most frequent element from the stack.

  Implement the FreqStack class:

    1. FreqStack() constructs an empty frequency stack.
    2. void push(int val) pushes an integer val onto the top of the stack.
    3. int pop() removes and returns the most frequent element in the stack.
    4. If there is a tie for the most frequent element, the element closest to the stack's top is removed and returned.


  Example 1:
    Input
      ["FreqStack", "push", "push", "push", "push", "push", "push", "pop", "pop", "pop", "pop"]
      [[], [5], [7], [5], [7], [4], [5], [], [], [], []]
    Output
      [null, null, null, null, null, null, null, 5, 7, 5, 4]
    Explanation
      FreqStack freqStack = new FreqStack();
      freqStack.push(5); // The stack is [5]
      freqStack.push(7); // The stack is [5,7]
      freqStack.push(5); // The stack is [5,7,5]
      freqStack.push(7); // The stack is [5,7,5,7]
      freqStack.push(4); // The stack is [5,7,5,7,4]
      freqStack.push(5); // The stack is [5,7,5,7,4,5]
      freqStack.pop();   // return 5, as 5 is the most frequent. The stack becomes [5,7,5,7,4].
      freqStack.pop();   // return 7, as 5 and 7 is the most frequent, but 7 is closest to the top. The stack becomes [5,7,5,4].
      freqStack.pop();   // return 5, as 5 is the most frequent. The stack becomes [5,7,4].
      freqStack.pop();   // return 4, as 4, 5 and 7 is the most frequent, but 4 is closest to the top. The stack becomes [5,7].

  Constraints:
    1. 0 <= val <= 10^9
    2. At most 2 * 10^4 calls will be made to push and pop.
    3. It is guaranteed that there will be at least one element in the stack before calling pop.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-frequency-stack
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Priority Queue(Heap Sort)

type FreqStack struct {
	cnt map[int]int
	idx int
	h   MaxHeap
}

func Constructor() FreqStack {
	return FreqStack{
		cnt: map[int]int{},
		idx: -1,
		h:   MaxHeap{},
	}
}

func (this *FreqStack) Push(val int) {
	this.cnt[val]++
	this.idx++
	heap.Push(&this.h, [3]int{val, this.cnt[val], this.idx})
}

func (this *FreqStack) Pop() int {
	top := heap.Pop(&this.h).([3]int)
	this.cnt[top[0]]--
	return top[0]
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */

type MaxHeap [][3]int // [val, freq, lastIdx]

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if h[i][1] != h[j][1] {
		return h[i][1] > h[j][1]
	}
	return h[i][2] > h[j][2]
}
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.([3]int)) }
func (h *MaxHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
