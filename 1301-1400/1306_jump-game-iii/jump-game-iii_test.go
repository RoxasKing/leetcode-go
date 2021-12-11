package main

import "testing"

func Test_canReach(t *testing.T) {
	type args struct {
		arr   []int
		start int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{4, 2, 3, 0, 3, 1, 2}, 5}, true},
		{"2", args{[]int{4, 2, 3, 0, 3, 1, 2}, 0}, true},
		{"3", args{[]int{3, 0, 2, 1, 2}, 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canReach(tt.args.arr, tt.args.start); got != tt.want {
				t.Errorf("canReach() = %v, want %v", got, tt.want)
			}
		})
	}
}
