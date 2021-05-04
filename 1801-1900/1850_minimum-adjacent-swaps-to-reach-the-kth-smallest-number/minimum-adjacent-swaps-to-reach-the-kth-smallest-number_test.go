package main

import "testing"

func Test_getMinSwaps(t *testing.T) {
	type args struct {
		num string
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"5489355142", 4}, 2},
		{"2", args{"11112", 4}, 4},
		{"3", args{"00123", 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinSwaps(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("getMinSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
