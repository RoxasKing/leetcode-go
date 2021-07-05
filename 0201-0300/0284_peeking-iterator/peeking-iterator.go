package main

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
