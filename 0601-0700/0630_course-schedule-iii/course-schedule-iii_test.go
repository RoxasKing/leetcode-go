package main

import "testing"

func Test_scheduleCourse(t *testing.T) {
	type args struct {
		courses [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}}, 3},
		{"2", args{[][]int{{1, 2}}}, 1},
		{"3", args{[][]int{{3, 2}, {4, 3}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scheduleCourse(tt.args.courses); got != tt.want {
				t.Errorf("scheduleCourse() = %v, want %v", got, tt.want)
			}
		})
	}
}
