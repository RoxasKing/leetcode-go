package main

/*
  给你一个嵌套的整型列表。请你设计一个迭代器，使其能够遍历这个整型列表中的所有整数。
  列表中的每一项或者为一个整数，或者是另一个列表。其中列表的元素也可能是整数或是其他列表。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/flatten-nested-list-iterator
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
	iter  *NestedIterator
	list  []*NestedInteger
	index int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	iter := &NestedIterator{list: nestedList}
	if len(nestedList) != 0 && !nestedList[0].IsInteger() {
		iter.iter = Constructor(nestedList[0].GetList())
	}
	return iter
}

func (this *NestedIterator) Next() int {
	if !this.list[this.index].IsInteger() {
		return this.iter.Next()
	}
	res := this.list[this.index].GetInteger()
	this.MakeNextIter()
	return res
}

func (this *NestedIterator) HasNext() bool {
	if this.index == len(this.list) {
		return false
	}
	if this.list[this.index].IsInteger() {
		return true
	}
	if !this.iter.HasNext() {
		this.MakeNextIter()
		return this.HasNext()
	}
	return true
}

func (this *NestedIterator) MakeNextIter() {
	this.index++
	if this.index < len(this.list) && !this.list[this.index].IsInteger() {
		this.iter = Constructor(this.list[this.index].GetList())
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
	if this.IsInteger() {
		return this.value
	}
	return -1 << 31
}

func (n *NestedInteger) SetInteger(value int) {
	n.isInteger = true
	n.value = value
	n.list = nil
}

func (this *NestedInteger) Add(elem NestedInteger) {
	if this.isInteger {
		this.isInteger = false
		this.value = -1 << 31
	}
	this.list = append(this.list, &elem)
}

func (this NestedInteger) GetList() []*NestedInteger {
	if this.IsInteger() {
		return nil
	}
	return this.list
}
