package main

import "testing"

func Test_kthToLast(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: &ListNode{Val: 5},
						},
					},
				},
			}, 2},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthToLast(tt.args.head, tt.args.k); got != tt.want {
				t.Errorf("kthToLast() = %v, want %v", got, tt.want)
			}
		})
	}
}
