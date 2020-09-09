package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{
				[]int{2, 3, 5},
				8,
			},
			[][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			"2",
			args{
				[]int{2, 3, 6, 7},
				7,
			},
			[][]int{
				{2, 2, 3},
				{7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum(tt.args.candidates, tt.args.target)
			sort.Slice(got, func(i, j int) bool {
				var index int
				for got[i][index] == got[j][index] {
					index++
				}
				return got[i][index] < got[j][index]
			})
			sort.Slice(tt.want, func(i, j int) bool {
				var index int
				for tt.want[i][index] == tt.want[j][index] {
					index++
				}
				return tt.want[i][index] < tt.want[j][index]
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{
				[]int{2, 3, 5},
				8,
			},
			[][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			"2",
			args{
				[]int{2, 3, 6, 7},
				7,
			},
			[][]int{
				{2, 2, 3},
				{7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum2(tt.args.candidates, tt.args.target)
			sort.Slice(got, func(i, j int) bool {
				var index int
				for got[i][index] == got[j][index] {
					index++
				}
				return got[i][index] < got[j][index]
			})
			sort.Slice(tt.want, func(i, j int) bool {
				var index int
				for tt.want[i][index] == tt.want[j][index] {
					index++
				}
				return tt.want[i][index] < tt.want[j][index]
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
