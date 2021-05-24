package main

import "testing"

func Test_minSpeedOnTime(t *testing.T) {
	type args struct {
		dist []int
		hour float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3, 2}, 6.0}, 1},
		{"2", args{[]int{1, 3, 2}, 2.7}, 3},
		{"3", args{[]int{1, 3, 2}, 1.9}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSpeedOnTime(tt.args.dist, tt.args.hour); got != tt.want {
				t.Errorf("minSpeedOnTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
