package main

// Difficulty:
// Medium

// Tags:
// Stack

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
	stk []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	for l, r := 0, len(nestedList)-1; l < r; l, r = l+1, r-1 {
		nestedList[l], nestedList[r] = nestedList[r], nestedList[l]
	}
	return &NestedIterator{nestedList}
}

func (this *NestedIterator) Next() int {
	this.help()
	i := len(this.stk) - 1
	out := this.stk[i].GetInteger()
	this.stk = this.stk[:i]
	return out
}

func (this *NestedIterator) HasNext() bool {
	this.help()
	return len(this.stk) > 0
}

func (this *NestedIterator) help() {
	for i := len(this.stk) - 1; i >= 0 && !this.stk[i].IsInteger(); i = len(this.stk) - 1 {
		nl := this.stk[i].GetList()
		this.stk = this.stk[:i]
		for l, r := 0, len(nl)-1; l < r; l, r = l+1, r-1 {
			nl[l], nl[r] = nl[r], nl[l]
		}
		this.stk = append(this.stk, nl...)
	}
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
