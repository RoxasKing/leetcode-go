package main

import "testing"

func Test_reformatNumber(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"1-23-45 6"}, "123-456"},
		{"2", args{"123 4-567"}, "123-45-67"},
		{"3", args{"123 4-5678"}, "123-456-78"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reformatNumber(tt.args.number); got != tt.want {
				t.Errorf("reformatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
