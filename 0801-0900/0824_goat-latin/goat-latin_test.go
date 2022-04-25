package main

import "testing"

func Test_toGoatLatin(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"I speak Goat Latin"}, "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"},
		{"2", args{"The quick brown fox jumped over the lazy dog"}, "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toGoatLatin(tt.args.sentence); got != tt.want {
				t.Errorf("toGoatLatin() = %v, want %v", got, tt.want)
			}
		})
	}
}
