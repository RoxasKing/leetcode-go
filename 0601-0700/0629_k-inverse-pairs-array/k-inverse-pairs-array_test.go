package main

import "testing"

func Test_kInversePairs(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{2, 0}, 1},
		{"2", args{2, 1}, 1},
		{"3", args{3, 0}, 1},
		{"4", args{3, 1}, 2},
		{"5", args{3, 2}, 2},
		{"6", args{3, 3}, 1},
		{"7", args{4, 1}, 3},
		{"8", args{4, 2}, 5},
		{"9", args{4, 3}, 6},
		{"10", args{4, 4}, 5},
		{"11", args{5, 1}, 4},
		{"12", args{5, 2}, 9},
		{"13", args{5, 3}, 15},
		{"14", args{5, 4}, 20},
		{"15", args{5, 5}, 22},
		{"16", args{6, 1}, 5},
		{"17", args{6, 2}, 14},
		{"18", args{6, 3}, 29},
		{"19", args{6, 4}, 49},
		{"20", args{6, 5}, 71},
		{"21", args{6, 6}, 90},
		{"22", args{6, 7}, 101},
		{"23", args{1000, 1000}, 663677020},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kInversePairs(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("kInversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
