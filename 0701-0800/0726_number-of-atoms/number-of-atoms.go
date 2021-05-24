package main

import (
	"sort"
	"strconv"
	"strings"
)

/*
  Given a chemical formula (given as a string), return the count of each atom.

  The atomic element always starts with an uppercase character, then zero or more lowercase letters, representing the name.

  One or more digits representing that element's count may follow if the count is greater than 1. If the count is 1, no digits will follow. For example, H2O and H2O2 are possible, but H1O2 is impossible.

  Two formulas concatenated together to produce another formula. For example, H2O2He3Mg4 is also a formula.

  A formula placed in parentheses, and a count (optionally added) is also a formula. For example, (H2O2) and (H2O2)3 are formulas.

  Given a formula, return the count of all elements as a string in the following form: the first name (in sorted order), followed by its count (if that count is more than 1), followed by the second name (in sorted order), followed by its count (if that count is more than 1), and so on.

  Example 1:
    Input: formula = "H2O"
    Output: "H2O"
    Explanation: The count of elements are {'H': 2, 'O': 1}.

  Example 2:
    Input: formula = "Mg(OH)2"
    Output: "H2MgO2"
    Explanation: The count of elements are {'H': 2, 'Mg': 1, 'O': 2}.

  Example 3:
    Input: formula = "K4(ON(SO3)2)2"
    Output: "K4N2O14S4"
    Explanation: The count of elements are {'K': 4, 'N': 2, 'O': 14, 'S': 4}.

  Example 4:
    Input: formula = "Be32"
    Output: "Be32"

  Constraints:
    1. 1 <= formula.length <= 1000
    2. formula consists of English letters, digits, '(', and ')'.
    3. formula is always valid.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/number-of-atoms
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Stack

func countOfAtoms(formula string) string {
	n := len(formula)
	as := AtomStack{}
	is := IntStack{}
	cnt, prev := 0, byte(' ')

	for i := 0; i < n; i++ {
		letter := formula[i]
		if isNumber(letter) {
			cnt = cnt*10 + int(letter-'0')
			continue
		} else if cnt > 0 {
			if prev == ')' {
				start := is.Pop()
				for j := start; j < as.Len(); j++ {
					as[j].freq *= cnt
				}
			} else {
				as.Top().freq *= cnt
			}
			cnt = 0
		}

		switch {
		case letter == '(':
			is.Push(as.Len())
		case isUpperCaseLetter(letter):
			if i+1 < n && isLowerCaseLetter(formula[i+1]) {
				as.Push(&Atom{name: formula[i : i+2], freq: 1})
				i++
			} else {
				as.Push(&Atom{name: formula[i : i+1], freq: 1})
			}
		}
		prev = letter
	}

	if cnt > 0 {
		if prev == ')' {
			start := is.Pop()
			for j := start; j < as.Len(); j++ {
				as[j].freq *= cnt
			}
		} else {
			as.Top().freq *= cnt
		}
	}

	freq := map[string]int{}
	for _, atom := range as {
		freq[atom.name] += atom.freq
	}

	strs := make([]string, 0, len(freq))
	for k, v := range freq {
		if v > 1 {
			k += strconv.Itoa(v)
		}
		strs = append(strs, k)
	}

	sort.Strings(strs)

	return strings.Join(strs, "")
}

func isUpperCaseLetter(letter byte) bool {
	return 'A' <= letter && letter <= 'Z'
}

func isLowerCaseLetter(letter byte) bool {
	return 'a' <= letter && letter <= 'z'
}

func isNumber(letter byte) bool {
	return '0' <= letter && letter <= '9'
}

type Atom struct {
	name string
	freq int
}

type AtomStack []*Atom

func (s AtomStack) Len() int      { return len(s) }
func (s AtomStack) Top() *Atom    { return s[s.Len()-1] }
func (s *AtomStack) Push(x *Atom) { *s = append(*s, x) }
func (s *AtomStack) Pop() *Atom {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}

type IntStack []int

func (s IntStack) Len() int    { return len(s) }
func (s IntStack) Top() int    { return s[s.Len()-1] }
func (s *IntStack) Push(x int) { *s = append(*s, x) }
func (s *IntStack) Pop() int {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}
