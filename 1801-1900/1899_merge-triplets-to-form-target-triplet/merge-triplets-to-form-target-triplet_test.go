package main

import "testing"

func Test_mergeTriplets(t *testing.T) {
	type args struct {
		triplets [][]int
		target   []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, []int{2, 7, 5}}, true},
		{"2", args{[][]int{{1, 3, 4}, {2, 5, 8}}, []int{2, 5, 8}}, true},
		{"3", args{[][]int{{2, 5, 3}, {2, 3, 4}, {1, 2, 5}, {5, 2, 3}}, []int{5, 5, 5}}, true},
		{"4", args{[][]int{{3, 4, 5}, {4, 5, 6}}, []int{3, 2, 5}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTriplets(tt.args.triplets, tt.args.target); got != tt.want {
				t.Errorf("mergeTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
