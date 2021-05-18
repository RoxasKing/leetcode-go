package main

import "testing"

func Test_countTriplets(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 3, 1, 6, 7}}, 4},
		{"2", args{[]int{1, 1, 1, 1, 1}}, 10},
		{"3", args{[]int{2, 3}}, 0},
		{"4", args{[]int{1, 3, 5, 7, 9}}, 3},
		{"5", args{[]int{7, 11, 12, 9, 5, 2, 7, 17, 22}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTriplets(tt.args.arr); got != tt.want {
				t.Errorf("countTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
