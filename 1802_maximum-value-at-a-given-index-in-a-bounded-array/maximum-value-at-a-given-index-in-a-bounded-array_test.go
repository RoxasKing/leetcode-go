package main

import "testing"

func Test_maxValue(t *testing.T) {
	type args struct {
		n      int
		index  int
		maxSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{4, 2, 6}, 2},
		{"2", args{6, 1, 10}, 3},
		{"3", args{4, 0, 4}, 1},
		{"4", args{7, 5, 30}, 6},
		{"5", args{2, 0, 173967479}, 86983740},
		{"6", args{1, 0, 1000}, 1000},
		{"7", args{2, 0, 1000}, 500},
		{"8", args{2, 1, 1000}, 500},
		{"9", args{2, 1, 1001}, 501},
		{"10", args{2, 1, 1003}, 502},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValue(tt.args.n, tt.args.index, tt.args.maxSum); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
