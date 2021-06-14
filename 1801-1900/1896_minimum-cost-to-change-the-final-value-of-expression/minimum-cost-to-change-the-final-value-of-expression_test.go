package main

import "testing"

func Test_minOperationsToFlip(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"1&(0|1)"}, 1},
		{"2", args{"(0&0)&(0&0&0)"}, 3},
		{"3", args{"(0|(1|0&1))"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperationsToFlip(tt.args.expression); got != tt.want {
				t.Errorf("minOperationsToFlip() = %v, want %v", got, tt.want)
			}
		})
	}
}
