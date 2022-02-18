package main

import "testing"

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  3,
							Next: &ListNode{Val: 4},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := []*ListNode{}
			for n := tt.args.head; n != nil; n = n.Next {
				arr = append(arr, n)
			}
			for i := 0; i+1 < len(arr); i += 2 {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
			got := swapPairs(tt.args.head)
			for n := got; n != nil; n = n.Next {
				if n != arr[0] {
					t.Errorf("swapPairs() test failed!")
					return
				}
				arr = arr[1:]
			}
		})
	}
}
