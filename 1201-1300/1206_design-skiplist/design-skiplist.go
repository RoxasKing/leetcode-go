package main

import "math/rand"

// Difficulty:
// Hard

// Tags:
// Design
// Linked List

const maxLvl = 32
const pFactor = 0.25

func randomLvl() int {
	lvl := 1
	for ; lvl < maxLvl && rand.Float64() < pFactor; lvl++ {
	}
	return lvl
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type slNode struct {
	val  int
	next []*slNode
}

type Skiplist struct {
	head *slNode
	lvl  int
}

func Constructor() Skiplist {
	return Skiplist{&slNode{-1, make([]*slNode, maxLvl)}, 0}
}

func (this *Skiplist) Search(target int) bool {
	t := this.head
	for i := this.lvl - 1; i >= 0; i-- {
		for ; t.next[i] != nil && t.next[i].val < target; t = t.next[i] {
		}
	}
	t = t.next[0]
	return t != nil && t.val == target
}

func (this *Skiplist) Add(num int) {
	upd := make([]*slNode, maxLvl)
	for i := range upd {
		upd[i] = this.head
	}
	t := this.head
	for i := this.lvl - 1; i >= 0; i-- {
		for ; t.next[i] != nil && t.next[i].val < num; t = t.next[i] {
		}
		upd[i] = t
	}
	lvl := randomLvl()
	this.lvl = max(this.lvl, lvl)
	newNode := &slNode{num, make([]*slNode, lvl)}
	for i, node := range upd[:lvl] {
		newNode.next[i] = node.next[i]
		node.next[i] = newNode
	}
}

func (this *Skiplist) Erase(num int) bool {
	upd := make([]*slNode, maxLvl)
	t := this.head
	for i := this.lvl - 1; i >= 0; i-- {
		for ; t.next[i] != nil && t.next[i].val < num; t = t.next[i] {
		}
		upd[i] = t
	}
	t = t.next[0]
	if t == nil || t.val != num {
		return false
	}
	for i := 0; i < this.lvl && upd[i].next[i] == t; i++ {
		upd[i].next[i] = t.next[i]
	}
	for ; this.lvl > 1 && this.head.next[this.lvl-1] == nil; this.lvl-- {
	}
	return true
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
