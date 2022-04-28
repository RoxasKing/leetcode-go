package main

// Difficulty:
// Medium

// Tags:
// Design

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
	val  int
	iter *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{-1, iter}
}

func (this *PeekingIterator) hasNext() bool {
	return this.val != -1 || this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.val != -1 {
		o := this.val
		this.val = -1
		return o
	}
	return this.iter.next()
}

func (this *PeekingIterator) peek() int {
	if this.val == -1 {
		this.val = this.iter.next()
	}
	return this.val
}

// ------------------------------------------------------------------------------------------------

type Iterator struct {
	a []int
	i int
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return this.i < len(this.a)
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	o := this.a[this.i]
	this.i++
	return o
}
