package main

import "testing"

func Test_maxTurbulenceSize(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{9, 4, 2, 10, 7, 8, 8, 1, 9}}, 5},
		{"2", args{[]int{4, 8, 12, 16}}, 2},
		{"3", args{[]int{100}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTurbulenceSize(tt.args.A); got != tt.want {
				t.Errorf("maxTurbulenceSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
