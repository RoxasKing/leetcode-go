package main

import (
	"reflect"
	"testing"
)

func Test_closestRoom(t *testing.T) {
	type args struct {
		rooms   [][]int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[][]int{{2, 2}, {1, 2}, {3, 2}}, [][]int{{3, 1}, {3, 3}, {5, 2}}}, []int{3, -1, 3}},
		{"2", args{[][]int{{1, 4}, {2, 3}, {3, 5}, {4, 1}, {5, 2}}, [][]int{{2, 3}, {2, 4}, {2, 5}}}, []int{2, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestRoom(tt.args.rooms, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("closestRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}
