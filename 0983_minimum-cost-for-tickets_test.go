package leetcode

import (
	"testing"
)

func Test_mincostTickets(t *testing.T) {
	type args struct {
		days  []int
		costs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[]int{1, 4, 6, 7, 8, 20},
				[]int{2, 7, 15},
			},
			11,
		},
		{
			"2",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
				[]int{2, 7, 15},
			},
			17,
		},
		{
			"3",
			args{
				[]int{6, 8, 9, 18, 20, 21, 23, 25},
				[]int{2, 10, 41},
			},
			16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mincostTickets(tt.args.days, tt.args.costs); got != tt.want {
				t.Errorf("mincostTickets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mincostTickets2(t *testing.T) {
	type args struct {
		days  []int
		costs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[]int{1, 4, 6, 7, 8, 20},
				[]int{2, 7, 15},
			},
			11,
		},
		{
			"2",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
				[]int{2, 7, 15},
			},
			17,
		},
		{
			"3",
			args{
				[]int{6, 8, 9, 18, 20, 21, 23, 25},
				[]int{2, 10, 41},
			},
			16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mincostTickets2(tt.args.days, tt.args.costs); got != tt.want {
				t.Errorf("mincostTickets2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mincostTickets3(t *testing.T) {
	type args struct {
		days  []int
		costs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[]int{1, 4, 6, 7, 8, 20},
				[]int{2, 7, 15},
			},
			11,
		},
		{
			"2",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
				[]int{2, 7, 15},
			},
			17,
		},
		{
			"3",
			args{
				[]int{6, 8, 9, 18, 20, 21, 23, 25},
				[]int{2, 10, 41},
			},
			16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mincostTickets3(tt.args.days, tt.args.costs); got != tt.want {
				t.Errorf("mincostTickets3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mincostTickets4(t *testing.T) {
	type args struct {
		days  []int
		costs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[]int{1, 4, 6, 7, 8, 20},
				[]int{2, 7, 15},
			},
			11,
		},
		{
			"2",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
				[]int{2, 7, 15},
			},
			17,
		},
		{
			"3",
			args{
				[]int{6, 8, 9, 18, 20, 21, 23, 25},
				[]int{2, 10, 41},
			},
			16,
		},
		{
			"4",
			args{
				[]int{2, 3, 4, 6, 8, 12, 14, 18, 19, 26, 27, 28},
				[]int{2, 9, 31},
			},
			23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mincostTickets4(tt.args.days, tt.args.costs); got != tt.want {
				t.Errorf("mincostTickets4() = %v, want %v", got, tt.want)
			}
		})
	}
}
