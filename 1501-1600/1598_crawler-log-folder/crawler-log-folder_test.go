package main

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		logs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"d1/", "d2/", "../", "d21/", "./"}}, 2},
		{"2", args{[]string{"d1/", "d2/", "./", "d3/", "../", "d31/"}}, 3},
		{"3", args{[]string{"d1/", "../", "../", "../"}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.logs); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
