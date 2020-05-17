package leetcode

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{2, [][]int{{1, 0}}}, []int{0, 1}},
		{"2", args{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}}, []int{0, 2, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findOrder2(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{2, [][]int{{1, 0}}}, []int{0, 1}},
		{"2", args{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}}, []int{0, 1, 2, 3}},
		{"3", args{3, [][]int{{1, 0}, {1, 2}, {0, 1}}}, nil},
		{"4", args{4, [][]int{{1, 0}, {2, 0}, {1, 3}, {2, 3}}}, []int{0, 3, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder2(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder2() = %v, want %v", got, tt.want)
			}
		})
	}
}
