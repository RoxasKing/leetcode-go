package main

import "testing"

func Test_getDecimalValue(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{&ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}, 5},
		{"2", args{&ListNode{Val: 0}}, 0},
		{"2", args{&ListNode{Val: 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDecimalValue(tt.args.head); got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
