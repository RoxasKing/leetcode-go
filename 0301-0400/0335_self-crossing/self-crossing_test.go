package main

import "testing"

func Test_isSelfCrossing(t *testing.T) {
	type args struct {
		distance []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{2, 1, 1, 2}}, true},
		{"2", args{[]int{1, 2, 3, 4}}, false},
		{"3", args{[]int{1, 1, 1, 1}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSelfCrossing(tt.args.distance); got != tt.want {
				t.Errorf("isSelfCrossing() = %v, want %v", got, tt.want)
			}
		})
	}
}
