package main

import "testing"

func Test_compareVersion(t *testing.T) {
	type args struct {
		version1 string
		version2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"0.1", "1.1"}, -1},
		{"2", args{"1.0.1", "1"}, 1},
		{"3", args{"7.5.2.4", "7.5.3"}, -1},
		{"4", args{"1.01", "1.001"}, 0},
		{"5", args{"1.0", "1.0.0"}, 0},
		{"6", args{"1.2", "1.10"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareVersion(tt.args.version1, tt.args.version2); got != tt.want {
				t.Errorf("compareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
