package main

import "testing"

func Test_leastOpsExpressTarget(t *testing.T) {
	type args struct {
		x      int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{3, 19}, 5},
		{"2", args{5, 501}, 8},
		{"3", args{100, 100000000}, 3},
		{"4", args{5, 109}, 7},
		{"5", args{5, 50}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastOpsExpressTarget(tt.args.x, tt.args.target); got != tt.want {
				t.Errorf("leastOpsExpressTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
