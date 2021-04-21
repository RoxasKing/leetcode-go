package main

import (
	"reflect"
	"testing"
)

func Test_allCellsDistOrder(t *testing.T) {
	type args struct {
		R  int
		C  int
		r0 int
		c0 int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{1, 2, 0, 0}, [][]int{{0, 0}, {0, 1}}},
		{"2", args{2, 2, 0, 1}, [][]int{{0, 1}, {0, 0}, {1, 1}, {1, 0}}},
		{"3", args{2, 3, 1, 2}, [][]int{{1, 2}, {0, 2}, {1, 1}, {0, 1}, {1, 0}, {0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allCellsDistOrder(tt.args.R, tt.args.C, tt.args.r0, tt.args.c0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allCellsDistOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allCellsDistOrder2(t *testing.T) {
	type args struct {
		R  int
		C  int
		r0 int
		c0 int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{1, 2, 0, 0}, [][]int{{0, 0}, {0, 1}}},
		{"2", args{2, 2, 0, 1}, [][]int{{0, 1}, {0, 0}, {1, 1}, {1, 0}}},
		{"3", args{2, 3, 1, 2}, [][]int{{1, 2}, {0, 2}, {1, 1}, {0, 1}, {1, 0}, {0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allCellsDistOrder2(tt.args.R, tt.args.C, tt.args.r0, tt.args.c0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allCellsDistOrder2() = %v, want %v", got, tt.want)
			}
		})
	}
}
