package main

import "testing"

func Test_expectNumber(t *testing.T) {
	type args struct {
		scores []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3}}, 3},
		{"2", args{[]int{1, 1}}, 1},
		{"3", args{[]int{1, 1, 2}}, 2},
		{"4", args{[]int{1, 1, 2, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expectNumber(tt.args.scores); got != tt.want {
				t.Errorf("expectNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
