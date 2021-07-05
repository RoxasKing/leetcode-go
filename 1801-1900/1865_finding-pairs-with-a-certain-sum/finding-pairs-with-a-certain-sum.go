package main

// Tags:
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
