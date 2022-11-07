package main

import (
	"reflect"
	"testing"
)

func Test_bestCoordinate(t *testing.T) {
	type args struct {
		towers [][]int
		radius int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[][]int{{1, 2, 5}, {2, 1, 7}, {3, 1, 9}}, 2}, []int{2, 1}},
		{"2", args{[][]int{{23, 11, 21}}, 9}, []int{23, 11}},
		{"3", args{[][]int{{1, 2, 13}, {2, 1, 7}, {0, 1, 9}}, 2}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestCoordinate(tt.args.towers, tt.args.radius); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bestCoordinate() = %v, want %v", got, tt.want)
			}
		})
	}
}
