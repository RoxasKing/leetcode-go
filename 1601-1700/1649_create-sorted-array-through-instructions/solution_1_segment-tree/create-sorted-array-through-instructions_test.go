package main

import "testing"

func Test_createSortedArray(t *testing.T) {
	type args struct {
		instructions []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 5, 6, 2}}, 1},
		{"2", args{[]int{1, 2, 3, 6, 5, 4}}, 3},
		{"3", args{[]int{1, 3, 3, 3, 2, 4, 2, 1, 2}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createSortedArray(tt.args.instructions); got != tt.want {
				t.Errorf("createSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
