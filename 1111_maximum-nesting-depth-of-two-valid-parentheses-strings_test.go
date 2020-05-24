package leetcode

import (
	"reflect"
	"testing"
)

func Test_maxDepthexp1afterSplit(t *testing.T) {
	type args struct {
		seq string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{"(()())"}, []int{1, 0, 0, 0, 0, 1}},
		{"2", args{"()(())()"}, []int{1, 1, 1, 0, 0, 1, 1, 1}},
		{"3", args{"((()))"}, []int{1, 0, 1, 1, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepthexp1afterSplit(tt.args.seq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxDepthexp1afterSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxDepthexp1afterSplit2(t *testing.T) {
	type args struct {
		seq string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{"(()())"}, []int{1, 0, 0, 0, 0, 1}},
		{"2", args{"()(())()"}, []int{1, 1, 1, 0, 0, 1, 1, 1}},
		{"3", args{"((()))"}, []int{1, 0, 1, 1, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepthexp1afterSplit2(tt.args.seq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxDepthexp1afterSplit2() = %v, want %v", got, tt.want)
			}
		})
	}
}
