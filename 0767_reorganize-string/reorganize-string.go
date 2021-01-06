package main

import (
	"container/heap"
)

/*
  给定一个字符串S，检查是否能重新排布其中的字母，使得两相邻的字符不同。
  若可行，输出任意可行的结果。若不可行，返回空字符串。

  示例 1:
    输入: S = "aab"
    输出: "aba"

  示例 2:
    输入: S = "aaab"
    输出: ""

  注意:
    S 只包含小写字母并且长度在[1, 500]区间内。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/reorganize-string
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Priority Queue(Heap Sort)
func reorganizeString(S string) string {
	n := len(S)
	if n <= 1 {
		return S
	}

	counts := [26]int{}
	for _, c := range S {
		index := int(c - 'a')
		if counts[index]++; counts[index] > (n+1)>>1 {
			return ""
		}
	}

	pq := priorityQueue{}

	for i, count := range counts {
		if count > 0 {
			heap.Push(&pq, &charCount{
				count: count,
				char:  byte(i) + 'a',
			})
		}
	}

	chars := make([]byte, 0, n)

	for pq.Len() > 1 {
		a, b := heap.Pop(&pq).(*charCount), heap.Pop(&pq).(*charCount)
		chars = append(chars, a.char, b.char)
		if a.count--; a.count > 0 {
			heap.Push(&pq, a)
		}
		if b.count--; b.count > 0 {
			heap.Push(&pq, b)
		}
	}

	if pq.Len() > 0 {
		chars = append(chars, pq.Pop().(*charCount).char)
	}

	return string(chars)
}

type charCount struct {
	count int
	char  byte
}

type priorityQueue []*charCount

func (p priorityQueue) Len() int            { return len(p) }
func (p priorityQueue) Less(i, j int) bool  { return p[i].count > p[j].count }
func (p priorityQueue) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *priorityQueue) Push(x interface{}) { *p = append(*p, x.(*charCount)) }
func (p *priorityQueue) Pop() interface{} {
	last := len(*p) - 1
	out := (*p)[last]
	*p = (*p)[:last]
	return out
}
