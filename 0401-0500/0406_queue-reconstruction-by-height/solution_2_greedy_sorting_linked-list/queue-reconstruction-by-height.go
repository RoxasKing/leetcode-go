package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Greedy
// Sorting
// Linked List

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	head := &Node{P: []int{-1, 0}}
	var count int
	for _, p := range people {
		head.Insert(p)
		count++
	}
	out := make([][]int, 0, count)
	for n := head.Next; n != nil; n = n.Next {
		out = append(out, n.P)
	}
	return out
}

type Node struct {
	P    []int
	Next *Node
}

func (n *Node) Insert(p []int) {
	var count int
	for count != p[1] {
		n = n.Next
		count++
	}
	next := n.Next
	n.Next = &Node{P: p}
	n.Next.Next = next
}
