package main

import "strconv"

// Dfficulty:
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
	n := len(s)
	stk := []NestedInteger{}
	top := func() int { return len(stk) - 1 }
	for c, i := 0, 0; i < n; i++ {
		switch s[i] {
		case ',':
			continue
		case '[':
			c++
			stk = append(stk, NestedInteger{})
			continue
		case ']':
			c--
		default:
			j := i
			for ; j+1 < n && ('0' <= s[j+1] && s[j+1] <= '9' || s[j+1] == '-'); j++ {
			}
			v, _ := strconv.Atoi(s[i : j+1])
			e := NestedInteger{}
			e.SetInteger(v)
			stk = append(stk, e)
			i = j
		}
		if c == 0 || len(stk) == 1 {
			continue
		}
		e := stk[top()]
		stk = stk[:top()]
		stk[top()].Add(e)
	}
	return &stk[0]
}

type NestedInteger struct {
	isInt bool
	value int
	nlist []*NestedInteger
}

func (n NestedInteger) IsInteger() bool       { return n.isInt }
func (n NestedInteger) GetInteger() int       { return n.value }
func (n *NestedInteger) SetInteger(value int) { n.isInt = true; n.value = value; n.nlist = nil }
func (n *NestedInteger) Add(elem NestedInteger) {
	n.isInt = false
	n.value = -1 << 31
	n.nlist = append(n.nlist, &elem)
}
func (n NestedInteger) GetList() []*NestedInteger { return n.nlist }
