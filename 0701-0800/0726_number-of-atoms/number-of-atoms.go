package main

import (
	"sort"
	"strconv"
	"strings"
)

// Tags:
// Stack

func countOfAtoms(formula string) string {
	atoms := AtomStack{}
	idxs := IntStack{}
	n := len(formula)
	for i := 0; i < n; i++ {
		letter := formula[i]
		switch {
		case letter == '(':
			idxs.Push(atoms.Len())
		case letter == ')':
			mul, j := 0, i
			for ; j+1 < n && isNumber(formula[j+1]); j++ {
				mul = mul*10 + int(formula[j+1]-'0')
			}
			i = j
			if mul > 0 {
				for k := idxs.Pop(); k < atoms.Len(); k++ {
					atoms[k].freq *= mul
				}
			}
		case isNumber(letter):
			mul, j := int(letter-'0'), i
			for ; j+1 < n && isNumber(formula[j+1]); j++ {
				mul = mul*10 + int(formula[j+1]-'0')
			}
			atoms.Top().freq *= mul
			i = j
		default:
			j := i
			for ; j+1 < n && isLowerCaseLetter(formula[j+1]); j++ {
			}
			atoms.Push(&Atom{name: formula[i : j+1], freq: 1})
			i = j
		}
	}

	freq := map[string]int{}
	for _, atom := range atoms {
		freq[atom.name] += atom.freq
	}
	arr := make([]string, 0, len(freq))
	for name := range freq {
		arr = append(arr, name)
	}
	sort.Strings(arr)

	var sb strings.Builder
	for _, name := range arr {
		sb.WriteString(name)
		if freq[name] > 1 {
			sb.WriteString(strconv.Itoa(freq[name]))
		}
	}
	return sb.String()
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
