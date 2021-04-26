package main

/*
  Design a data structure that is initialized with a list of different words. Provided a string, you should determine if you can change exactly one character in this string to match any word in the data structure.

  Implement the MagicDictionary class:
    1. MagicDictionary() Initializes the object.
    2. void buildDict(String[] dictionary) Sets the data structure with an array of distinct strings dictionary.
    3. bool search(String searchWord) Returns true if you can change exactly one character in searchWord to match any string in the data structure, otherwise returns false.


  Example 1:
    Input
      ["MagicDictionary", "buildDict", "search", "search", "search", "search"]
      [[], [["hello", "leetcode"]], ["hello"], ["hhllo"], ["hell"], ["leetcoded"]]
    Output
      [null, null, false, true, false, false]
    Explanation
      MagicDictionary magicDictionary = new MagicDictionary();
      magicDictionary.buildDict(["hello", "leetcode"]);
      magicDictionary.search("hello"); // return False
      magicDictionary.search("hhllo"); // We can change the second 'h' to 'e' to match "hello" so we return True
      magicDictionary.search("hell"); // return False
      magicDictionary.search("leetcoded"); // return False

  Constraints:
    1. 1 <= dictionary.length <= 100
    2. 1 <= dictionary[i].length <= 100
    3. dictionary[i] consists of only lower-case English letters.
    4. All the strings in dictionary are distinct.
    5. 1 <= searchWord.length <= 100
    6. searchWord consists of only lower-case English letters.
    7. buildDict will be called only once before search.
    8. At most 100 calls will be made to search.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/implement-magic-dictionary
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
