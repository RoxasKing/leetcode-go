package main

import "strings"

// Difficulty:
// Medium

func validIPAddress(queryIP string) string {
	if a := strings.Split(queryIP, "."); len(a) == 4 {
		ok := true
		for _, s := range a {
			if len(s) < 1 || 3 < len(s) {
				ok = false
				break
			}
			for _, c := range s {
				if c < '0' || '9' < c {
					ok = false
					break
				}
			}
			if len(s) > 1 && s[0] == '0' {
				ok = false
				break
			}
			if len(s) == 3 && (s[0] > '2' || s[0] == '2' && (s[1] > '5' || s[1] == '5' && s[2] > '5')) {
				ok = false
				break
			}
			if !ok {
				break
			}
		}
		if ok {
			return "IPv4"
		}
	}
	if a := strings.Split(queryIP, ":"); len(a) == 8 {
		ok := true
		for _, s := range a {
			if len(s) < 1 || 4 < len(s) {
				ok = false
				break
			}
			for _, c := range s {
				if !('0' <= c && c <= '9' || 'a' <= c && c <= 'f' || 'A' <= c && c <= 'F') {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
		}
		if ok {
			return "IPv6"
		}
	}
	return "Neither"
}
