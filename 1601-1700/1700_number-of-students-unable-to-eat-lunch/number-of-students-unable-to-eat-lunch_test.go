package main

import "testing"

func Test_countStudents(t *testing.T) {
	type args struct {
		students   []int
		sandwiches []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 1, 0, 0}, []int{0, 1, 0, 1}}, 0},
		{"2", args{[]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countStudents(tt.args.students, tt.args.sandwiches); got != tt.want {
				t.Errorf("countStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
