package main

// Tags:
// Recursion

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	list []*NestedInteger
	idx  int
	iter *NestedIterator
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{list: nestedList}
}

func (this *NestedIterator) Next() int {
	if !this.list[this.idx].IsInteger() {
		return this.iter.Next()
	}
	out := this.list[this.idx].GetInteger()
	this.idx++
	return out
}

func (this *NestedIterator) HasNext() bool {
START:
	if this.idx == len(this.list) {
		return false
	}

	if this.list[this.idx].IsInteger() {
		return true
	}

	if this.iter == nil {
		this.iter = Constructor(this.list[this.idx].GetList())
	}

	if this.iter.HasNext() {
		return true
	}

	this.idx++
	this.iter = nil
	goto START
}

type NestedInteger struct {
	isInteger bool
	value     int
	list      []*NestedInteger
}

func (this NestedInteger) IsInteger() bool {
	return this.isInteger
}

func (this NestedInteger) GetInteger() int {
	return this.value
}

func (n *NestedInteger) SetInteger(value int) {
	n.isInteger = true
	n.value = value
	n.list = nil
}

func (this *NestedInteger) Add(elem NestedInteger) {
	this.isInteger = false
	this.value = -1 << 31
	this.list = append(this.list, &elem)
}

func (this NestedInteger) GetList() []*NestedInteger {
	return this.list
}
