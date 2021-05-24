package main

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"<DIV>This is the first line <![CDATA[<div>]]></DIV>"}, true},
		{"2", args{"<DIV>>>  ![cdata[]] <![CDATA[<div>]>]]>]]>>]</DIV>"}, true},
		{"3", args{"<A>  <B> </A>   </B>"}, false},
		{"4", args{"<DIV>  div tag is not closed  <DIV>"}, false},
		{"5", args{"<DIV>  unmatched <  </DIV>"}, false},
		{"6", args{"<DIV> closed tags with invalid tag name  <b>123</b> </DIV>"}, false},
		{"7", args{"<DIV> unmatched tags with invalid tag name  </1234567890> and <CDATA[[]]>  </DIV>"}, false},
		{"8", args{"<DIV>  unmatched start tag <B>  and unmatched end tag </C>  </DIV>"}, false},
		{"9", args{"<a><![CDATA[wahaha]]]><![CDATA[]> wahaha]]></a>"}, false},
		{"10", args{"<A></A><B></B>"}, false},
		{"11", args{"<TRUE><![CDATA[wahaha]]]><![CDATA[]> wahaha]]></TRUE>"}, true},
		{"12", args{"<A><A></A></A>"}, true},
		{"13", args{"<A></A>"}, true},
		{"14", args{"<![CDATA[wahaha]]]><![CDATA[]> wahaha]]>"}, false},
		{"15", args{"<></>"}, false},
		{"16", args{"<AAAAAAAAAA></AAAAAAAAAA>"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.code); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
