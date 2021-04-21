package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{&ListNode{Val: 1, Next: &ListNode{Val: 2}}}, false},
		{
			"2",
			args{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  2,
							Next: &ListNode{Val: 1},
						},
					},
				},
			},
			true,
		},
		{"3", args{&ListNode{Val: 1}}, true},
		{"4", args{&ListNode{Val: 0, Next: &ListNode{Val: 0}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
