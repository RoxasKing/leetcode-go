package main

import "strconv"

/*
  Given a string s represents the serialization of a nested list, implement a parser to deserialize it and return the deserialized NestedInteger.

  Each element is either an integer or a list whose elements may also be integers or other lists.

  Example 1:
    Input: s = "324"
    Output: 324
    Explanation: You should return a NestedInteger object which contains a single integer 324.
    Example 2:
      Input: s = "[123,[456,[789]]]"
      Output: [123,[456,[789]]]
      Explanation: Return a NestedInteger object containing a nested list with 2 elements:
        1. An integer containing value 123.
        2. A nested list containing two elements:
            i.  An integer containing value 456.
            ii. A nested list with one element:
                 a. An integer containing value 789

  Constraints:
    1. 1 <= s.length <= 5 * 10^4
    2. s consists of digits, square brackets "[]", negative sign '-', and commas ','.
    3. s is the serialization of valid NestedInteger.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/mini-parser
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

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
