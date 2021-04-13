package main

import (
	"testing"
)

func Test_numPermsDISequence(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"DID"}, 5},
		{"2", args{"IDID"}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numPermsDISequence(tt.args.S); got != tt.want {
				t.Errorf("numPermsDISequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numPermsDISequence2(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"DID"}, 5},
		{"2", args{"IDID"}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numPermsDISequence2(tt.args.S); got != tt.want {
				t.Errorf("numPermsDISequence2() = %v, want %v", got, tt.want)
			}
		})
	}
}
