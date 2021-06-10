package main

import (
	"reflect"
	"testing"
)

func Test_waysToFillArray(t *testing.T) {
	type args struct {
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[][]int{{2, 6}, {5, 1}, {73, 660}}}, []int{4, 1, 50734910}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToFillArray(tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("waysToFillArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
