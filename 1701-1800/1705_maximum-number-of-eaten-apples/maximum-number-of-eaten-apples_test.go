package main

import "testing"

func Test_eatenApples(t *testing.T) {
	type args struct {
		apples []int
		days   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 5, 2}, []int{3, 2, 1, 4, 2}}, 7},
		{"2", args{[]int{3, 0, 0, 0, 0, 2}, []int{3, 0, 0, 0, 0, 2}}, 5},
		{"3", args{[]int{2, 1, 1, 4, 5}, []int{10, 10, 6, 4, 2}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eatenApples(tt.args.apples, tt.args.days); got != tt.want {
				t.Errorf("eatenApples() = %v, want %v", got, tt.want)
			}
		})
	}
}
