package main

import (
	"testing"
)

func Test_findCheapestPrice(t *testing.T) {
	type args struct {
		n       int
		flights [][]int
		src     int
		dst     int
		K       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				3,
				[][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
				0,
				2,
				1,
			},
			200,
		},
		{
			"2",
			args{
				3,
				[][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
				0,
				2,
				0,
			},
			500,
		},
		{
			"3",
			args{
				4,
				[][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}, {2, 3, 50}, {1, 3, 100}},
				0,
				3,
				2,
			},
			200,
		},
		{
			"4",
			args{
				5,
				[][]int{{4, 1, 1}, {1, 2, 3}, {0, 3, 2}, {0, 4, 10}, {3, 1, 1}, {1, 4, 3}},
				2,
				1,
				1,
			},
			-1,
		},
		{
			"5",
			args{
				10,
				[][]int{{3, 4, 4}, {2, 5, 6}, {4, 7, 10}, {9, 6, 5}, {7, 4, 4}, {6, 2, 10}, {6, 8, 6}, {7, 9, 4}, {1, 5, 4}, {1, 0, 4}, {9, 7, 3}, {7, 0, 5}, {6, 5, 8}, {1, 7, 6}, {4, 0, 9}, {5, 9, 1}, {8, 7, 3}, {1, 2, 6}, {4, 1, 5}, {5, 2, 4}, {1, 9, 1}, {7, 8, 10}, {0, 4, 2}, {7, 2, 8}},
				6,
				0,
				7,
			},
			14,
		},
		{
			"6",
			args{
				4,
				[][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}},
				0,
				3,
				1,
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCheapestPrice(tt.args.n, tt.args.flights, tt.args.src, tt.args.dst, tt.args.K); got != tt.want {
				t.Errorf("findCheapestPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findCheapestPrice2(t *testing.T) {
	type args struct {
		n       int
		flights [][]int
		src     int
		dst     int
		K       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				3,
				[][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
				0,
				2,
				1,
			},
			200,
		},
		{
			"2",
			args{
				3,
				[][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
				0,
				2,
				0,
			},
			500,
		},
		{
			"3",
			args{
				4,
				[][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}, {2, 3, 50}, {1, 3, 100}},
				0,
				3,
				2,
			},
			200,
		},
		{
			"4",
			args{
				5,
				[][]int{{4, 1, 1}, {1, 2, 3}, {0, 3, 2}, {0, 4, 10}, {3, 1, 1}, {1, 4, 3}},
				2,
				1,
				1,
			},
			-1,
		},
		{
			"5",
			args{
				10,
				[][]int{{3, 4, 4}, {2, 5, 6}, {4, 7, 10}, {9, 6, 5}, {7, 4, 4}, {6, 2, 10}, {6, 8, 6}, {7, 9, 4}, {1, 5, 4}, {1, 0, 4}, {9, 7, 3}, {7, 0, 5}, {6, 5, 8}, {1, 7, 6}, {4, 0, 9}, {5, 9, 1}, {8, 7, 3}, {1, 2, 6}, {4, 1, 5}, {5, 2, 4}, {1, 9, 1}, {7, 8, 10}, {0, 4, 2}, {7, 2, 8}},
				6,
				0,
				7,
			},
			14,
		},
		{
			"6",
			args{
				4,
				[][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}},
				0,
				3,
				1,
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCheapestPrice2(tt.args.n, tt.args.flights, tt.args.src, tt.args.dst, tt.args.K); got != tt.want {
				t.Errorf("findCheapestPrice2() = %v, want %v", got, tt.want)
			}
		})
	}
}
