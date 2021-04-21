package main

import "testing"

func Test_isMonotonic(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{1, 2, 2, 3}}, true},
		{"2", args{[]int{6, 5, 4, 4}}, true},
		{"3", args{[]int{1, 3, 2}}, false},
		{"4", args{[]int{1, 2, 4, 5}}, true},
		{"5", args{[]int{1, 1, 1}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMonotonic(tt.args.A); got != tt.want {
				t.Errorf("isMonotonic() = %v, want %v", got, tt.want)
			}
		})
	}
}
