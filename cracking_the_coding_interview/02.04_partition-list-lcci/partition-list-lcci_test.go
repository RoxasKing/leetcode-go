package main

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"1",
			args{&ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 8,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 10,
								Next: &ListNode{
									Val:  2,
									Next: &ListNode{Val: 1},
								},
							},
						},
					},
				},
			}, 5},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 8,
								Next: &ListNode{
									Val:  5,
									Next: &ListNode{Val: 10},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
