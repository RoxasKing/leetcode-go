package codinginterviews

import (
	"testing"
)

func Test_movingCount(t *testing.T) {
	type args struct {
		m int
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{2, 3, 1},
			3,
		},
		{
			"2",
			args{3, 1, 0},
			1,
		},
		{
			"3",
			args{16, 8, 4},
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movingCount(tt.args.m, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("movingCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_movingCount2(t *testing.T) {
	type args struct {
		m int
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{2, 3, 1},
			3,
		},
		{
			"2",
			args{3, 1, 0},
			1,
		},
		{
			"3",
			args{16, 8, 4},
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movingCount2(tt.args.m, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("movingCount2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_movingCount3(t *testing.T) {
	type args struct {
		m int
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{2, 3, 1},
			3,
		},
		{
			"2",
			args{3, 1, 0},
			1,
		},
		{
			"3",
			args{16, 8, 4},
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movingCount3(tt.args.m, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("movingCount3() = %v, want %v", got, tt.want)
			}
		})
	}
}
