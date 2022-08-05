package main

import "testing"

func Test_validSquare(t *testing.T) {
	type args struct {
		p1 []int
		p2 []int
		p3 []int
		p4 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 1}}, true},
		{"2", args{[]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 12}}, false},
		{"3", args{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validSquare(tt.args.p1, tt.args.p2, tt.args.p3, tt.args.p4); got != tt.want {
				t.Errorf("validSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
