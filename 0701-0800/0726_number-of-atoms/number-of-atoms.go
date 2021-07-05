package main

import (
	"sort"
	"strconv"
	"strings"
)

// Tags:
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
