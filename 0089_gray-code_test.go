package leetcode

import (
	"reflect"
	"testing"
)

func Test_grayCode(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{0}, []int{0}},
		{"", args{2}, []int{0, 1, 3, 2}},
		{"", args{3}, []int{0, 1, 3, 2, 6, 7, 5, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grayCode(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("grayCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
