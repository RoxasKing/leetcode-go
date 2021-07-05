package main

import (
	"strconv"
	"strings"
)

// Tags:
// Stack

func exclusiveTime(n int, logs []string) []int {
	out := make([]int, n)
	fs := FuncStack{}
	for _, log := range logs {
		strs := strings.Split(log, ":")
		id, _ := strconv.Atoi(strs[0])
		tm, _ := strconv.Atoi(strs[2])
		if strs[1] == "start" {
			fs.Push(&Func{id: id, st: tm, ex: 0})
			continue
		}
		fu := fs.Pop()
		out[id] += tm + 1 - fu.st - fu.ex
		if fs.Len() > 0 {
			fs.Top().ex += tm + 1 - fu.st
		}
	}
	return out
}

type Func struct {
	id, st, ex int
}

type FuncStack []*Func

func (s FuncStack) Len() int      { return len(s) }
func (s FuncStack) Top() *Func    { return s[s.Len()-1] }
func (s *FuncStack) Push(x *Func) { *s = append(*s, x) }
func (s *FuncStack) Pop() *Func {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}
