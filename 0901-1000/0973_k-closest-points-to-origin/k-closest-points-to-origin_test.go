package main

import (
	"reflect"
	"testing"
)

func Test_kClosest(t *testing.T) {
	type args struct {
		points [][]int
		K      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[][]int{{1, 3}, {-2, 2}}, 1}, [][]int{{-2, 2}}},
		{"2", args{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2}, [][]int{{3, 3}, {-2, 4}}},
		{"3", args{[][]int{{-5, 4}, {6, -5}, {4, 6}}, 2}, [][]int{{-5, 4}, {4, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kClosest(tt.args.points, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kClosest2(t *testing.T) {
	type args struct {
		points [][]int
		K      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[][]int{{1, 3}, {-2, 2}}, 1}, [][]int{{-2, 2}}},
		{"2", args{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2}, [][]int{{3, 3}, {-2, 4}}},
		{"3", args{[][]int{{-5, 4}, {6, -5}, {4, 6}}, 2}, [][]int{{-5, 4}, {4, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kClosest2(tt.args.points, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosest2() = %v, want %v", got, tt.want)
			}
		})
	}
}
