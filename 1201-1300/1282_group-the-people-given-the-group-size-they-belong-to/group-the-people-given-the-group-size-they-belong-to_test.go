package main

import (
	"reflect"
	"testing"
)

func Test_groupThePeople(t *testing.T) {
	type args struct {
		groupSizes []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[]int{3, 3, 3, 3, 3, 1, 3}}, [][]int{{5}, {0, 1, 2}, {3, 4, 6}}},
		{"2", args{[]int{2, 1, 3, 3, 3, 2}}, [][]int{{1}, {0, 5}, {2, 3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupThePeople(tt.args.groupSizes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupThePeople() = %v, want %v", got, tt.want)
			}
		})
	}
}
