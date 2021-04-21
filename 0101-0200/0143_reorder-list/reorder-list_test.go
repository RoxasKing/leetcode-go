package main

import "testing"

func Test_reorderList(t *testing.T) {
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
							Val: 3,
							Next: &ListNode{
								Val: 4,
							},
						},
					},
				},
			},
		},
		{
			"2",
			args{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
		},
		{
			"3",
			args{nil},
		},
		{
			"4",
			args{&ListNode{Val: 1}},
		},
		{
			"5",
			args{&ListNode{Val: 1, Next: &ListNode{Val: 2}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var order []*ListNode
			head := tt.args.head
			for n := head; n != nil; n = n.Next {
				order = append(order, n)
			}
			if len(order) > 1 {
				m := (len(order) - 1) >> 1
				o1, o2 := make([]*ListNode, m+1), make([]*ListNode, len(order)-1-m)
				copy(o1, order[:m+1])
				copy(o2, order[m+1:])
				for i := 0; i < len(o2)>>1; i++ {
					o2[i], o2[len(o2)-1-i] = o2[len(o2)-1-i], o2[i]
				}
				size := len(order)
				order = make([]*ListNode, 0, size)
				for i := 0; i < size; i++ {
					if len(o1) != 0 {
						order = append(order, o1[0])
						o1 = o1[1:]
					}
					i++
					if len(o2) != 0 {
						order = append(order, o2[0])
						o2 = o2[1:]
					}
				}
			}
			reorderList(tt.args.head)
			for n := tt.args.head; n != nil; n = n.Next {
				if len(order) == 0 || n != order[0] {
					t.Errorf("reorderList() test failed")
					return
				}
				order = order[1:]
			}
			if len(order) != 0 {
				t.Errorf("reorderList() test failed")
			}
		})
	}
}
