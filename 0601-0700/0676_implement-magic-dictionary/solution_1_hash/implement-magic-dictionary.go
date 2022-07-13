package main

// Difficulty:
// Medium

// Tags:
// Hash

type MagicDictionary struct {
	h map[int][]string
}

func Constructor() MagicDictionary {
	return MagicDictionary{h: map[int][]string{}}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, w := range dictionary {
		this.h[len(w)] = append(this.h[len(w)], w)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	if _, ok := this.h[len(searchWord)]; !ok {
		return false
	}
	for _, w := range this.h[len(searchWord)] {
		c := 0
		for i := range w {
			if w[i] == searchWord[i] {
				continue
			}
			if c++; c > 1 {
				break
			}
		}
		if c == 1 {
			return true
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
