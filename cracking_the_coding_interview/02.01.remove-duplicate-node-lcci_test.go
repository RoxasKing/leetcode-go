package crackingthecodingintervew

import (
	"reflect"
	"testing"
)

func Test_removeDuplicateNodes(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
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
							Val: 3,
							Next: &ListNode{
								Val:  2,
								Next: &ListNode{Val: 1},
							},
						},
					},
				},
			}},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: &ListNode{Val: 3},
				},
			},
		},
		{"2", args{&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  1,
						Next: &ListNode{Val: 2},
					},
				},
			},
		}},
			&ListNode{
				Val:  1,
				Next: &ListNode{Val: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicateNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicateNodes2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
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
							Val: 3,
							Next: &ListNode{
								Val:  2,
								Next: &ListNode{Val: 1},
							},
						},
					},
				},
			}},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: &ListNode{Val: 3},
				},
			},
		},
		{"2", args{&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  1,
						Next: &ListNode{Val: 2},
					},
				},
			},
		}},
			&ListNode{
				Val:  1,
				Next: &ListNode{Val: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateNodes2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicateNodes2() = %v, want %v", got, tt.want)
			}
		})
	}
}
