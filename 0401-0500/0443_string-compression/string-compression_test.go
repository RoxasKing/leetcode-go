package main

import "testing"

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}}, 6},
		{"2", args{[]byte{'a'}}, 1},
		{"3", args{[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}}, 4},
		{"4", args{[]byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.chars); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
		})
	}
}
