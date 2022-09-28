package main

import "strings"

// Difficulty:
// Medium

// Tags:
// Hash

func findDuplicate(paths []string) [][]string {
	h := map[string][]string{}
	for _, path := range paths {
		a := strings.Split(path, " ")
		for _, s := range a[1:] {
			ss := strings.Split(s, "(")
			key := ss[1][:len(ss[1])-1]
			h[key] = append(h[key], a[0]+"/"+ss[0])
		}
	}
	o := [][]string{}
	for _, list := range h {
		if len(list) > 1 {
			o = append(o, list)
		}
	}
	return o
}
