package main

import "fmt"

func findSubstring2(s string, words []string) []int {
	var out []int


	return out
}

func findSubstring(s string, words []string) []int {
	var out []int
	sizes := len(s)
	sizew := len(words)
	if sizes == 0 || sizew == 0 {
		return nil
	}
	sizee := len(words[0])
	sizea := sizew * sizee
	dict := make(map[string]int)
	for _, elem := range words {
		dict[elem]++
	}
	var flg bool
	for i := 0; i < sizes-sizea+1; i++ {
		for j := i; j < i+sizea-sizee+1; j = j + sizee {
			if _, ok := dict[s[j:j+sizee]]; ok {
				dict[s[j:j+sizee]]--
			}
		}
		for _, v := range dict {
			if v != 0 {
				flg = true
				break
			}
		}
		dict = make(map[string]int)
		for _, elem := range words {
			dict[elem]++
		}
		if !flg {
			out = append(out, i)
		} else {
			flg = false
		}
	}
	return out
}

func main() {
	//s := "barfoothefoobarman"
	//words := []string{"foo", "bar"}
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	//s := ""
	//words := []string{}
	fmt.Println(findSubstring(s, words))
}
