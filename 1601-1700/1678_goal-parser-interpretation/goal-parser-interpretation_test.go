package main

import "testing"

func Test_interpret(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"G()(al)"}, "Goal"},
		{"2", args{"G()()()()(al)"}, "Gooooal"},
		{"3", args{"(al)G(al)()()G"}, "alGalooG"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := interpret(tt.args.command); got != tt.want {
				t.Errorf("interpret() = %v, want %v", got, tt.want)
			}
		})
	}
}
