package main

import "testing"

func Test_kthLargestValue(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{5, 2}, {1, 6}}, 1}, 7},
		{"2", args{[][]int{{5, 2}, {1, 6}}, 2}, 5},
		{"3", args{[][]int{{5, 2}, {1, 6}}, 3}, 4},
		{"4", args{[][]int{{5, 2}, {1, 6}}, 4}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargestValue(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthLargestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
