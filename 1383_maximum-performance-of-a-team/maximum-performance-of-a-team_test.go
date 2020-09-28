package main

import "testing"

func Test_maxPerformance(t *testing.T) {
	type args struct {
		n          int
		speed      []int
		efficiency []int
		k          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 2}, 60},
		{"2", args{6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 3}, 68},
		{"3", args{6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 4}, 72},
		{"4", args{3, []int{2, 8, 2}, []int{2, 7, 1}, 2}, 56},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPerformance(tt.args.n, tt.args.speed, tt.args.efficiency, tt.args.k); got != tt.want {
				t.Errorf("maxPerformance() = %v, want %v", got, tt.want)
			}
		})
	}
}
