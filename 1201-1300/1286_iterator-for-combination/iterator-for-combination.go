package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Design
// Iterator

type CombinationIterator struct {
	chs []byte
	idx []int
	ptr int
	flg bool
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	chs := []byte(characters)
	sort.Slice(chs, func(i, j int) bool { return chs[i] < chs[j] })
	idx := make([]int, combinationLength)
	for i := range idx {
		idx[i] = i
	}
	return CombinationIterator{chs: chs, idx: idx, ptr: combinationLength - 1, flg: true}
}

func (this *CombinationIterator) Next() string {
	if !this.flg {
		this.next()
	}
	this.flg = false
	out := make([]byte, len(this.idx))
	for i := range out {
		out[i] = this.chs[this.idx[i]]
	}
	return string(out)
}

func (this *CombinationIterator) HasNext() bool {
	if this.flg {
		return true
	}
	if this.next() {
		this.flg = true
		return true
	}
	return false
}

func (this *CombinationIterator) next() bool {
START:
	if this.ptr < 0 {
		return false
	}
	for i := this.ptr; i < len(this.idx); i++ {
		if i > this.ptr {
			this.idx[i] = this.idx[i-1] + 1
		} else {
			this.idx[i]++
		}
		if this.idx[i] == len(this.chs) {
			this.ptr--
			goto START
		}
	}
	this.ptr = len(this.idx) - 1
	return true
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
