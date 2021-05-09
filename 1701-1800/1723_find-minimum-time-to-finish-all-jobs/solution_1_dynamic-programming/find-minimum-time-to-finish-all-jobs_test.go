package main

import "testing"

func Test_minimumTimeRequired(t *testing.T) {
	type args struct {
		jobs []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 2, 3}, 3}, 3},
		{"2", args{[]int{1, 2, 4, 7, 8}, 2}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTimeRequired(tt.args.jobs, tt.args.k); got != tt.want {
				t.Errorf("minimumTimeRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}
