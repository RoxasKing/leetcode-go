package main

import "testing"

func Test_brokenCalc(t *testing.T) {
	type args struct {
		startValue int
		target     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{2, 3}, 2},
		{"2", args{5, 8}, 2},
		{"3", args{3, 10}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := brokenCalc(tt.args.startValue, tt.args.target); got != tt.want {
				t.Errorf("brokenCalc() = %v, want %v", got, tt.want)
			}
		})
	}
}
