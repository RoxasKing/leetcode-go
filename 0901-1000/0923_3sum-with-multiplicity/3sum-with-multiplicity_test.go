package main

import "testing"

func Test_threeSumMulti(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8}, 20},
		{"2", args{[]int{1, 1, 2, 2, 2, 2}, 5}, 12},
		{"3", args{[]int{0, 0, 0}, 0}, 1},
		{"4", args{[]int{3, 3, 1}, 7}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumMulti(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("threeSumMulti() = %v, want %v", got, tt.want)
			}
		})
	}
}
