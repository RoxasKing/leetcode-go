package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_accountsMerge(t *testing.T) {
	type args struct {
		accounts [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"1",
			args{[][]string{
				{"John", "johnsmith@mail.com", "john00@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
				{"Mary", "mary@mail.com"},
			}},
			[][]string{
				{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"Mary", "mary@mail.com"},
			},
		},
		{
			"2",
			args{[][]string{
				{"David", "David0@m.co", "David1@m.co"},
				{"David", "David3@m.co", "David4@m.co"},
				{"David", "David4@m.co", "David5@m.co"},
				{"David", "David2@m.co", "David3@m.co"},
				{"David", "David1@m.co", "David2@m.co"},
			}},
			[][]string{
				{"David", "David0@m.co", "David1@m.co", "David2@m.co", "David3@m.co", "David4@m.co", "David5@m.co"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := accountsMerge(tt.args.accounts)
			sortStrSlices(got)
			sortStrSlices(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("accountsMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sortStrSlices(strs [][]string) {
	sort.Slice(strs, func(i, j int) bool {
		for k := 0; k < Min(len(strs[i]), len(strs[j])); k++ {
			if strs[i][k] != strs[j][k] {
				return strs[i][k] < strs[j][k]
			}
		}
		return len(strs[i]) < len(strs[j])
	})
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
