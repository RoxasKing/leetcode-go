package main

import "testing"

func Test_findMaximizedCapital(t *testing.T) {
	type args struct {
		k       int
		w       int
		profits []int
		capital []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{2, 0, []int{1, 2, 3}, []int{0, 1, 1}}, 4},
		{"2", args{3, 0, []int{1, 2, 3}, []int{0, 1, 2}}, 6},
		{"3", args{1, 2, []int{1, 2, 3}, []int{1, 1, 2}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaximizedCapital(tt.args.k, tt.args.w, tt.args.profits, tt.args.capital); got != tt.want {
				t.Errorf("findMaximizedCapital() = %v, want %v", got, tt.want)
			}
		})
	}
}
