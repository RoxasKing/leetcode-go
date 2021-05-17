package main

/*
  You are given two integer arrays nums1 and nums2. You are tasked to implement a data structure that supports queries of two types:
    1. Add a positive integer to an element of a given index in the array nums2.
    2. Count the number of pairs (i, j) such that nums1[i] + nums2[j] equals a given value (0 <= i < nums1.length and 0 <= j < nums2.length).

  Implement the FindSumPairs class:
    1. FindSumPairs(int[] nums1, int[] nums2) Initializes the FindSumPairs object with two integer arrays nums1 and nums2.
    2. void add(int index, int val) Adds val to nums2[index], i.e., apply nums2[index] += val.
    3. int count(int tot) Returns the number of pairs (i, j) such that nums1[i] + nums2[j] == tot.

  Example 1:
    Input
      ["FindSumPairs", "count", "add", "count", "count", "add", "add", "count"]
      [[[1, 1, 2, 2, 2, 3], [1, 4, 5, 2, 5, 4]], [7], [3, 2], [8], [4], [0, 1], [1, 1], [7]]
    Output
      [null, 8, null, 2, 1, null, null, 11]
    Explanation
      FindSumPairs findSumPairs = new FindSumPairs([1, 1, 2, 2, 2, 3], [1, 4, 5, 2, 5, 4]);
      findSumPairs.count(7);  // return 8; pairs (2,2), (3,2), (4,2), (2,4), (3,4), (4,4) make 2 + 5 and pairs (5,1), (5,5) make 3 + 4
      findSumPairs.add(3, 2); // now nums2 = [1,4,5,4,5,4]
      findSumPairs.count(8);  // return 2; pairs (5,2), (5,4) make 3 + 5
      findSumPairs.count(4);  // return 1; pair (5,0) makes 3 + 1
      findSumPairs.add(0, 1); // now nums2 = [2,4,5,4,5,4]
      findSumPairs.add(1, 1); // now nums2 = [2,5,5,4,5,4]
      findSumPairs.count(7);  // return 11; pairs (2,1), (2,2), (2,4), (3,1), (3,2), (3,4), (4,1), (4,2), (4,4) make 2 + 5 and pairs (5,3), (5,5) make 3 + 4

  Constraints:
    1. 1 <= nums1.length <= 1000
    2. 1 <= nums2.length <= 10^5
    3. 1 <= nums1[i] <= 10^9
    4. 1 <= nums2[i] <= 10^5
    5. 0 <= index < nums2.length
    6. 1 <= val <= 10^5
    7. 1 <= tot <= 10^9
    8. At most 1000 calls are made to add and count each.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/finding-pairs-with-a-certain-sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Hash

type FindSumPairs struct {
	cnt1 map[int]int
	cnt2 map[int]int
	arr  []int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	cnt1 := map[int]int{}
	for _, num := range nums1 {
		cnt1[num]++
	}
	cnt2 := map[int]int{}
	for _, num := range nums2 {
		cnt2[num]++
	}
	return FindSumPairs{
		cnt1: cnt1,
		cnt2: cnt2,
		arr:  nums2,
	}
}

func (this *FindSumPairs) Add(index int, val int) {
	this.cnt2[this.arr[index]]--
	if this.cnt2[this.arr[index]] == 0 {
		delete(this.cnt2, this.arr[index])
	}
	this.arr[index] += val
	this.cnt2[this.arr[index]]++
}

func (this *FindSumPairs) Count(tot int) int {
	out := 0
	cnt1, cnt2 := this.cnt1, this.cnt2
	if len(cnt1) > len(cnt2) {
		cnt1, cnt2 = cnt2, cnt1
	}
	for num, freq := range cnt1 {
		out += freq * cnt2[tot-num]
	}
	return out
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
