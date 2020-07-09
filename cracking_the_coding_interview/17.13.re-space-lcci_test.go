package crackingthecodingintervew

import (
	"testing"
)

func Test_respace(t *testing.T) {
	type args struct {
		dictionary []string
		sentence   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[]string{"looked", "just", "like", "her", "brother"},
				"jesslookedjustliketimherbrother",
			},
			7,
		},
		{
			"2",
			args{
				[]string{"qqqqqqqqq", "qqqqq", "qq", "qqq", "qqqqqq", "qqqqqq", "q"},
				"qqqq",
			},
			0,
		},
		{
			"3",
			args{
				[]string{"pvoviipvpiop", "vvvivppioi", "vpivxiixoxpixpioiivoopvivoi", "opovvixov", "pvo", "vpoioiov", "v", "oixopvvpvxip", "xx", "ooppp", "iovxxx", "opxvvioiiivoixvioxox", "ipixixpioxvovoooivixpxoiooppoipiovxvoixvovpxxiixvpipovooxoppiippvxvvoipioopvpoipopppppoipv", "pvxxxoxpoo", "i", "ivvoxovvpv", "oopivoioixoopop", "ovvppoixvvixvviipxppxipixiooiivxxxxxp", "v", "xvooxpxxpiioipxixxopxovpiioxoxpxovioooppoiioiooovippoppopivpvoipxxxxvpvxvovpvpopvvxppvxvppx", "v", "ixvop", "opopioiv", "vvooi", "vi", "ov", "xvxooxopixxiioipixovvvovii", "vpiopipvooippixxioxxoxvpioivxoxxxixvxpiovovxiioivvopvppoixoiiiiivpvpxoppix", "vo", "oivov", "poxvvovixvxvxiioxoi", "pxoivivoioviippvovxivvxvpooxxxppxvxpiov", "iiio", "i", "vx", "xpvvxxviiiooxixpviopxv", "pxpp", "pppxxvvvoviviiooooipvvpvippxvoiovoiipppviovioixovpooxipvpoxxvxpoovixpxpoxvvooxvoopi", "vopv", "iv", "vxpv", "pivxpvipixxvvvvxxivpvipoxpivpiopiovxoooppoxpxipvvooipooovvoivppvoixixxpvipvxvxvoivx"},
				"pppxxvvvoviviiooooipvvpvippxvoiovoiipppviovioixovpooxipvpoxxvxpoovixpxpoxvvooxvoopivxpvxxxxpxpp",
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := respace(tt.args.dictionary, tt.args.sentence); got != tt.want {
				t.Errorf("respace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_respace2(t *testing.T) {
	type args struct {
		dictionary []string
		sentence   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[]string{"looked", "just", "like", "her", "brother"},
				"jesslookedjustliketimherbrother",
			},
			7,
		},
		{
			"2",
			args{
				[]string{"qqqqqqqqq", "qqqqq", "qq", "qqq", "qqqqqq", "qqqqqq", "q"},
				"qqqq",
			},
			0,
		},
		{
			"3",
			args{
				[]string{"pvoviipvpiop", "vvvivppioi", "vpivxiixoxpixpioiivoopvivoi", "opovvixov", "pvo", "vpoioiov", "v", "oixopvvpvxip", "xx", "ooppp", "iovxxx", "opxvvioiiivoixvioxox", "ipixixpioxvovoooivixpxoiooppoipiovxvoixvovpxxiixvpipovooxoppiippvxvvoipioopvpoipopppppoipv", "pvxxxoxpoo", "i", "ivvoxovvpv", "oopivoioixoopop", "ovvppoixvvixvviipxppxipixiooiivxxxxxp", "v", "xvooxpxxpiioipxixxopxovpiioxoxpxovioooppoiioiooovippoppopivpvoipxxxxvpvxvovpvpopvvxppvxvppx", "v", "ixvop", "opopioiv", "vvooi", "vi", "ov", "xvxooxopixxiioipixovvvovii", "vpiopipvooippixxioxxoxvpioivxoxxxixvxpiovovxiioivvopvppoixoiiiiivpvpxoppix", "vo", "oivov", "poxvvovixvxvxiioxoi", "pxoivivoioviippvovxivvxvpooxxxppxvxpiov", "iiio", "i", "vx", "xpvvxxviiiooxixpviopxv", "pxpp", "pppxxvvvoviviiooooipvvpvippxvoiovoiipppviovioixovpooxipvpoxxvxpoovixpxpoxvvooxvoopi", "vopv", "iv", "vxpv", "pivxpvipixxvvvvxxivpvipoxpivpiopiovxoooppoxpxipvvooipooovvoivppvoixixxpvipvxvxvoivx"},
				"pppxxvvvoviviiooooipvvpvippxvoiovoiipppviovioixovpooxipvpoxxvxpoovixpxpoxvvooxvoopivxpvxxxxpxpp",
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := respace2(tt.args.dictionary, tt.args.sentence); got != tt.want {
				t.Errorf("respace2() = %v, want %v", got, tt.want)
			}
		})
	}
}
