package main

import "testing"

func Test_numberOfWeeks(t *testing.T) {
	type args struct {
		milestones []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"1", args{[]int{1, 2, 3}}, 6},
		{"2", args{[]int{5, 2, 1}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWeeks(tt.args.milestones); got != tt.want {
				t.Errorf("numberOfWeeks() = %v, want %v", got, tt.want)
			}
		})
	}
}
