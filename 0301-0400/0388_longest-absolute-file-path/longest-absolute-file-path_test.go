package main

import "testing"

func Test_lengthLongestPath(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"}, 20},
		{"2", args{"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"}, 32},
		{"3", args{"a"}, 0},
		{"4", args{"file1.txt\nfile2.txt\nlongfile.txt"}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthLongestPath(tt.args.input); got != tt.want {
				t.Errorf("lengthLongestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
