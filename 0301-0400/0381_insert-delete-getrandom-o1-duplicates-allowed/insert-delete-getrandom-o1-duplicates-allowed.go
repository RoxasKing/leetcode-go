package main

import (
	"math/rand"
	"sort"
)

/*
  Implement the RandomizedCollection class:

    1. RandomizedCollection() Initializes the RandomizedCollection object.
    2. bool insert(int val) Inserts an item val into the multiset if not present. Returns true if the item was not present, false otherwise.
    3. bool remove(int val) Removes an item val from the multiset if present. Returns true if the item was present, false otherwise. Note that if val has multiple occurrences in the multiset, we only remove one of them.
    4. int getRandom() Returns a random element from the current multiset of elements (it's guaranteed that at least one element exists when this method is called). The probability of each element being returned is linearly related to the number of same values the multiset contains.

  Example 1:
    Input
      ["RandomizedCollection", "insert", "insert", "insert", "getRandom", "remove", "getRandom"]
      [[], [1], [1], [2], [], [1], []]
    Output
      [null, true, false, true, 2, true, 1]
    Explanation
      RandomizedCollection randomizedCollection = new RandomizedCollection();
      randomizedCollection.insert(1);   // return True. Inserts 1 to the collection. Returns true as the collection did not contain 1.
      randomizedCollection.insert(1);   // return False. Inserts another 1 to the collection. Returns false as the collection contained 1. Collection now contains [1,1].
      randomizedCollection.insert(2);   // return True. Inserts 2 to the collection, returns true. Collection now contains [1,1,2].
      randomizedCollection.getRandom(); // getRandom should return 1 with the probability 2/3, and returns 2 with the probability 1/3.
      randomizedCollection.remove(1);   // return True. Removes 1 from the collection, returns true. Collection now contains [1,2].
      randomizedCollection.getRandom(); // getRandom should return 1 and 2 both equally likely.

  Constraints:
    1. -2^31 <= val <= 2^31 - 1
    2. At most 10^5 calls will be made to insert, remove, and getRandom.
    3. There will be at least one element in the data structure when getRandom is called.

  Follow up: Could you implement the functions of the class with each function works in average O(1) time?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/insert-delete-getrandom-o1-duplicates-allowed
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Hash
type RandomizedCollection struct {
	idxs map[int][]int
	list []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		idxs: make(map[int][]int),
		list: []int{},
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	success := false
	if _, ok := this.idxs[val]; !ok {
		success = true
	}
	this.idxs[val] = append(this.idxs[val], len(this.list))
	this.list = append(this.list, val)
	return success
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	if arr, ok := this.idxs[val]; !ok || len(arr) == 0 {
		return false
	}
	idx := this.idxs[val][len(this.idxs[val])-1]
	this.idxs[val] = this.idxs[val][:len(this.idxs[val])-1]
	tail := len(this.list) - 1
	tailVal := this.list[tail]
	if tailVal == val {
		this.list = this.list[:tail]
		return true
	}
	this.list[tail], this.list[idx] = this.list[idx], this.list[tail]
	this.list = this.list[:tail]
	this.idxs[tailVal] = append(this.idxs[tailVal][:len(this.idxs[tailVal])-1], idx)
	sort.Ints(this.idxs[tailVal])
	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
