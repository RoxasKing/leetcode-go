package main

// Tags:
// Hash
type MagicDictionary struct {
	groupedByLen map[int][]string
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{
		groupedByLen: make(map[int][]string),
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, w := range dictionary {
		this.groupedByLen[len(w)] = append(this.groupedByLen[len(w)], w)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	for _, w := range this.groupedByLen[len(searchWord)] {
		if this.checkIfOneDiff(searchWord, w) {
			return true
		}
	}
	return false
}

func (this *MagicDictionary) checkIfOneDiff(w1, w2 string) bool {
	cnt := 0
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			cnt++
			if cnt > 1 {
				return false
			}
		}
	}
	return cnt == 1
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
