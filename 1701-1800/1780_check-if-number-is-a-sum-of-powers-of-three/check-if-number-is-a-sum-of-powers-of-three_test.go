package main

import "testing"

func Test_checkPowersOfThree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{12}, true},
		{"2", args{91}, true},
		{"3", args{21}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPowersOfThree(tt.args.n); got != tt.want {
				t.Errorf("checkPowersOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
