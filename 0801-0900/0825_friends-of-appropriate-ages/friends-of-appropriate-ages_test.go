package main

import "testing"

func Test_numFriendRequests(t *testing.T) {
	type args struct {
		ages []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{16, 16}}, 2},
		{"2", args{[]int{16, 17, 18}}, 2},
		{"3", args{[]int{20, 30, 100, 110, 120}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numFriendRequests(tt.args.ages); got != tt.want {
				t.Errorf("numFriendRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}
