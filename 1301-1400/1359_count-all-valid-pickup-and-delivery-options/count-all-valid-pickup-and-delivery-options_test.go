package main

import "testing"

func Test_countOrders(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1}, 1},
		{"2", args{2}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOrders(tt.args.n); got != tt.want {
				t.Errorf("countOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}
