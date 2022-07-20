package main

import "testing"

func Test_containVirus(t *testing.T) {
	type args struct {
		isInfected [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{"1", args{[][]int{{0, 1, 0, 0, 0, 0, 0, 1}, {0, 1, 0, 0, 0, 0, 0, 1}, {0, 0, 0, 0, 0, 0, 0, 1}, {0, 0, 0, 0, 0, 0, 0, 0}}}, 10},
		{"2", args{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}}, 4},
		{"3", args{[][]int{{1, 1, 1, 0, 0, 0, 0, 0, 0}, {1, 0, 1, 0, 1, 1, 1, 1, 1}, {1, 1, 1, 0, 0, 0, 0, 0, 0}}}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := containVirus(tt.args.isInfected); gotAns != tt.wantAns {
				t.Errorf("containVirus() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
