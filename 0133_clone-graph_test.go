package leetcode

import (
	"testing"
)

func Test_cloneGraph(t *testing.T) {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node1.Neighbors = append(node1.Neighbors, node2)
	node2.Neighbors = append(node2.Neighbors, node1)

	nodea := &Node{Val: 1}
	nodeb := &Node{Val: 2}
	nodec := &Node{Val: 3}
	noded := &Node{Val: 4}
	nodea.Neighbors = append(nodea.Neighbors, nodeb, noded)
	nodeb.Neighbors = append(nodeb.Neighbors, nodea, nodec)
	nodec.Neighbors = append(nodec.Neighbors, nodeb, noded)
	noded.Neighbors = append(noded.Neighbors, nodea, nodec)
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"1", args{nil}, nil},
		{"2", args{&Node{Val: 1}}, &Node{Val: 1}},
		{"3", args{node1}, node1},
		{"4", args{nodea}, nodea},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cloneGraph(tt.args.node)
			if tt.want == nil && got == nil {
				return
			}
			mark := make(map[int]bool)
			queue1, queue2 := []*Node{got}, []*Node{tt.want}
			for len(queue1) != 0 {
				N1, N2 := queue1[0], queue2[0]
				queue1, queue2 = queue1[1:], queue2[1:]
				if N1 == N2 || N1.Val != N2.Val || len(N1.Neighbors) != len(N2.Neighbors) {
					t.Errorf("cloneGraph() = %v, want %v", got, tt.want)
					return
				}
				if mark[N1.Val] {
					continue
				}
				for i := range N1.Neighbors {
					queue1 = append(queue1, N1.Neighbors[i])
					queue2 = append(queue2, N2.Neighbors[i])
				}
				mark[N1.Val] = true
			}
		})
	}
}
