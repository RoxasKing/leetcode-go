package main

import (
	"reflect"
	"testing"
)

func Test_imageSmoother(t *testing.T) {
	type args struct {
		img [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}},
			[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			"2",
			args{[][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}},
			[][]int{{137, 141, 137}, {141, 138, 141}, {137, 141, 137}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := imageSmoother(tt.args.img); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageSmoother() = %v, want %v", got, tt.want)
			}
		})
	}
}
