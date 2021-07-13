package main

import "testing"

func Test_sumGame(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"5023"}, false},
		{"2", args{"25??"}, true},
		{"3", args{"?3295???"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumGame(tt.args.num); got != tt.want {
				t.Errorf("sumGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
