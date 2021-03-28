package main

import "testing"

func Test_isValidSerialization(t *testing.T) {
	type args struct {
		preorder string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"9,3,4,#,#,1,#,#,2,#,6,#,#"}, true},
		{"2", args{"1,#"}, false},
		{"3", args{"9,#,#,1"}, false},
		{"4", args{"1,#,#"}, true},
		{"5", args{"1,2,3,#,#,#,#"}, true},
		{"6", args{"#"}, true},
		{"7", args{"#,2"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSerialization(tt.args.preorder); got != tt.want {
				t.Errorf("isValidSerialization() = %v, want %v", got, tt.want)
			}
		})
	}
}
