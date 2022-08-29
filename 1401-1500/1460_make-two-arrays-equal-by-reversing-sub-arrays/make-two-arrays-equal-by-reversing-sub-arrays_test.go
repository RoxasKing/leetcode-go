package main

import "testing"

func Test_canBeEqual(t *testing.T) {
	type args struct {
		target []int
		arr    []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{1, 2, 3, 4}, []int{2, 4, 1, 3}}, true},
		{"2", args{[]int{7}, []int{7}}, true},
		{"3", args{[]int{3, 7, 9}, []int{3, 7, 11}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeEqual(tt.args.target, tt.args.arr); got != tt.want {
				t.Errorf("canBeEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
