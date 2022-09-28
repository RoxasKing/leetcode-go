package main

import "testing"

func Test_concatenatedBinary(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1}, 1},
		{"2", args{3}, 27},
		{"3", args{12}, 505379714},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := concatenatedBinary(tt.args.n); got != tt.want {
				t.Errorf("concatenatedBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
