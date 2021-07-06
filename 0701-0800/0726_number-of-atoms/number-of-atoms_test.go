package main

import "testing"

func Test_countOfAtoms(t *testing.T) {
	type args struct {
		formula string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"H2O"}, "H2O"},
		{"2", args{"Mg(OH)2"}, "H2MgO2"},
		{"3", args{"K4(ON(SO3)2)2"}, "K4N2O14S4"},
		{"4", args{"Be32"}, "Be32"},
		{"5", args{"Mg(H2O)N"}, "H2MgNO"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfAtoms(tt.args.formula); got != tt.want {
				t.Errorf("countOfAtoms() = %v, want %v", got, tt.want)
			}
		})
	}
}
