package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}, 5}, 8},
		{"2", args{[]int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}, 11}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
