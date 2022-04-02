package main

import "testing"

func Test_strongPasswordChecker(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"a"}, 5},
		{"2", args{"aA1"}, 3},
		{"3", args{"1337C0d3"}, 0},
		{"4", args{"ABABABABABABABABABAB1"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strongPasswordChecker(tt.args.password); got != tt.want {
				t.Errorf("strongPasswordChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
