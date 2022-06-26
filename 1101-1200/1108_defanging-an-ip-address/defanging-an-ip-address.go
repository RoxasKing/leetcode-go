package main

import "strings"

// Difficulty:
// Easy

func defangIPaddr(address string) string {
	return strings.Join(strings.Split(address, "."), "[.]")
}
