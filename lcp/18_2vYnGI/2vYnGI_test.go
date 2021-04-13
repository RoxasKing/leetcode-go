package main

import "testing"

func Test_breakfastNumber(t *testing.T) {
	type args struct {
		staple []int
		drinks []int
		x      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{10, 20, 5}, []int{5, 5, 2}, 15}, 6},
		{"2", args{[]int{2, 1, 1}, []int{8, 9, 5, 1}, 9}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := breakfastNumber(tt.args.staple, tt.args.drinks, tt.args.x); got != tt.want {
				t.Errorf("breakfastNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
