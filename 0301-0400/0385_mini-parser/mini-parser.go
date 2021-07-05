package main

import "strconv"

// Stack

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
	stk := []*NestedInteger{}
	n := len(s)
	cnt := 0

	for i := 0; i < n; i++ {
		if s[i] == ',' {
			continue
		} else if s[i] == '[' {
			stk = append(stk, &NestedInteger{})
			cnt++
			continue
		} else if s[i] == ']' {
			cnt--
		} else {
			j := i
			for j+1 < n && ('0' <= s[j+1] && s[j+1] <= '9' || s[j+1] == '-') {
				j++
			}
			num, _ := strconv.Atoi(s[i : j+1])
			ni := &NestedInteger{}
			ni.SetInteger(num)
			stk = append(stk, ni)
			i = j
		}

		if cnt == 0 || len(stk) < 2 {
			continue
		}

		top := len(stk) - 1
		ni := stk[top]
		stk = stk[:top]
		pni := stk[top-1]
		pni.Add(*ni)
	}

	return stk[0]
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
