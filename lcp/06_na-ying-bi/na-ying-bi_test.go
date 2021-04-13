package main

import "testing"

func Test_minCount(t *testing.T) {
	type args struct {
		coins []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{4, 2, 1}}, 4},
		{"2", args{[]int{2, 3, 10}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCount(tt.args.coins); got != tt.want {
				t.Errorf("minCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
