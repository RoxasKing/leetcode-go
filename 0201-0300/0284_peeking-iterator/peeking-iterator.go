package main

/*
  Design an iterator that supports the peek operation on a list in addition to the hasNext and the next operations.

  Implement the PeekingIterator class:

    1. PeekingIterator(int[] nums) Initializes the object with the given integer array nums.
    2. int next() Returns the next element in the array and moves the pointer to the next element.
    3. bool hasNext() Returns true if there are still elements in the array.
    4. int peek() Returns the next element in the array without moving the pointer.

  Example 1:
    Input
      ["PeekingIterator", "next", "peek", "next", "next", "hasNext"]
      [[[1, 2, 3]], [], [], [], [], []]
    Output
      [null, 1, 2, 2, 3, false]
    Explanation
      PeekingIterator peekingIterator = new PeekingIterator([1, 2, 3]); // [1,2,3]
      peekingIterator.next();    // return 1, the pointer moves to the next element [1,2,3].
      peekingIterator.peek();    // return 2, the pointer does not move [1,2,3].
      peekingIterator.next();    // return 2, the pointer moves to the next element [1,2,3]
      peekingIterator.next();    // return 3, the pointer moves to the next element [1,2,3]
      peekingIterator.hasNext(); // return False

  Constraints:
    1. 1 <= nums.length <= 1000
    2. 1 <= nums[i] <= 1000
    3. All the calls to next and peek are valid.
    4. At most 1000 calls will be made to next, hasNext, and peek.

    Follow up: How would you extend your design to be generic and work with all types, not just integer?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/peeking-iterator
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
	iter *Iterator
	cur  int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter: iter, cur: -1}
}

func (this *PeekingIterator) hasNext() bool {
	return this.cur != -1 || this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.cur != -1 {
		out := this.cur
		this.cur = -1
		return out
	}
	return this.iter.next()
}

func (this *PeekingIterator) peek() int {
	if this.cur != -1 {
		return this.cur
	}
	this.cur = this.iter.next()
	return this.cur
}

// ------------------------------------------------------------------------------------------------

type Iterator struct {
	arr []int
	i   int
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return this.i < len(this.arr)
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	out := this.arr[this.i]
	this.i++
	return out
}
