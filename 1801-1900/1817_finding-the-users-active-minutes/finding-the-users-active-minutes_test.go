package main

import (
	"reflect"
	"testing"
)

func Test_findingUsersActiveMinutes(t *testing.T) {
	type args struct {
		logs [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}}, 5}, []int{0, 2, 0, 0, 0}},
		{"2", args{[][]int{{1, 1}, {2, 2}, {2, 3}}, 4}, []int{1, 1, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findingUsersActiveMinutes(tt.args.logs, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findingUsersActiveMinutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
