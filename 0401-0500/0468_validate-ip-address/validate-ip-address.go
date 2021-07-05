package main

import "strings"

func validIPAddress(IP string) string {
	if arr := strings.Split(IP, "."); len(arr) == 4 {
		if ok := validIPv4(arr); ok {
			return "IPv4"
		}
	}
	if arr := strings.Split(IP, ":"); len(arr) == 8 {
		if ok := validIPv6(arr); ok {
			return "IPv6"
		}
	}
	return "Neither"
}

func validIPv4(arr []string) bool {
	for _, n := range arr {
		if len(n) < 1 || 3 < len(n) {
			return false
		}
		for i := range n {
			if !('0' <= n[i] && n[i] <= '9') {
				return false
			}
		}
		if len(n) > 1 && n[0] == '0' {
			return false
		}
		if len(n) == 3 && (n[0] > '2' || n[0] == '2' && (n[1] > '5' || n[1] == '5' && n[2] > '5')) {
			return false
		}
	}
	return true
}

func validIPv6(arr []string) bool {
	for _, n := range arr {
		if len(n) < 1 || 4 < len(n) {
			return false
		}
		for i := range n {
			if !('0' <= n[i] && n[i] <= '9' || 'a' <= n[i] && n[i] <= 'f' || 'A' <= n[i] && n[i] <= 'F') {
				return false
			}
		}
	}
	return true
}
