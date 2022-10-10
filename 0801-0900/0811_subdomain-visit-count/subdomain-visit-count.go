package main

import (
	"strconv"
	"strings"
)

// Difficulty:
// Medium

// Tags:
// Hash

func subdomainVisits(cpdomains []string) []string {
	h := map[string]int{}
	for _, domain := range cpdomains {
		strs := strings.Split(domain, " ")
		cnt, _ := strconv.Atoi(strs[0])
		a := strings.Split(strs[1], ".")
		for i := len(a) - 1; i >= 0; i-- {
			h[strings.Join(a[i:], ".")] += cnt
		}
	}
	o := []string{}
	for str, cnt := range h {
		o = append(o, strconv.Itoa(cnt)+" "+str)
	}
	return o
}
