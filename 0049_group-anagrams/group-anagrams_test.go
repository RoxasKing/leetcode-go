package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"1",
			args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			[][]string{
				{"ate", "eat", "tea"},
				{"nat", "tan"},
				{"bat"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.args.strs)
			sort.Slice(got, func(i, j int) bool {
				return len(got[i]) > len(got[j])
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return len(tt.want[i]) > len(got[j])
			})
			for i := range got {
				sort.Slice(got[i], func(j, k int) bool {
					return got[i][j][0] < got[i][k][0]
				})
			}
			for i := range tt.want {
				sort.Slice(tt.want[i], func(j, k int) bool {
					return tt.want[i][j][0] < tt.want[i][k][0]
				})
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupAnagrams2(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"2",
			args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			[][]string{
				{"ate", "eat", "tea"},
				{"nat", "tan"},
				{"bat"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams2(tt.args.strs)
			sort.Slice(got, func(i, j int) bool {
				return len(got[i]) > len(got[j])
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return len(tt.want[i]) > len(got[j])
			})
			for i := range got {
				sort.Slice(got[i], func(j, k int) bool {
					return got[i][j][0] < got[i][k][0]
				})
			}
			for i := range tt.want {
				sort.Slice(tt.want[i], func(j, k int) bool {
					return tt.want[i][j][0] < tt.want[i][k][0]
				})
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams2() = %v, want %v", got, tt.want)
			}
		})
	}
}
