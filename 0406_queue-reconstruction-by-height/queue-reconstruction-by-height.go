package main

import (
	"sort"
)

/*
  假设有打乱顺序的一群人站成一个队列。 每个人由一个整数对(h, k)表示，其中h是这个人的身高，k是排在这个人前面且身高大于或等于h的人数。 编写一个算法来重建这个队列。

  注意：
    总人数少于1100人。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/queue-reconstruction-by-height
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Greedy Algorithm + Array
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	var out [][]int
	for _, p := range people {
		out = append(out, p)
		copy(out[p[1]+1:], out[p[1]:])
		out[p[1]] = p
	}
	return out
}

// Greedy Algorithm + LinkedList
func reconstructQueue2(people [][]int) [][]int {
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
