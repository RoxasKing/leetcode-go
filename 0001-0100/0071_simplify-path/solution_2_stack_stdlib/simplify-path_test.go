package main

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"/a/./b///../c/../././../d/..//../e/./f/./g/././//.//h///././/..///"}, "/e/f/g"},
		{"2", args{"/.."}, "/"},
		{"3", args{"/."}, "/"},
		{"4", args{"/home/"}, "/home"},
		{"5", args{"/../"}, "/"},
		{"6", args{"/home//foo/"}, "/home/foo"},
		{"7", args{"/a/./b/../../c/"}, "/c"},
		{"8", args{"/a/../../b/../c//.//"}, "/c"},
		{"9", args{"/a//b////c/d//././/.."}, "/a/b/c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
