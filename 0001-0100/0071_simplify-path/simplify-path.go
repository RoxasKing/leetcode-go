package main

import "strings"

func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	tmp := []string{}
	for _, dir := range dirs {
		if dir == "." || dir == "" {
			continue
		} else if dir != ".." {
			tmp = append(tmp, dir)
		} else if len(tmp) > 0 {
			tmp = tmp[:len(tmp)-1]
		}
	}
	return "/" + strings.Join(tmp, "/")
}
