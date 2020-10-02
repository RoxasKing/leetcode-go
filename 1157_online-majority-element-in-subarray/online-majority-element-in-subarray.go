package main

import "sort"

/*
  实现一个 MajorityChecker 的类，它应该具有下述几个 API：

  MajorityChecker(int[] arr) 会用给定的数组 arr 来构造一个 MajorityChecker 的实例。
  int query(int left, int right, int threshold) 有这么几个参数：
  0 <= left <= right < arr.length 表示数组 arr 的子数组的长度。
  2 * threshold > right - left + 1，也就是说阈值 threshold 始终比子序列长度的一半还要大。
  每次查询 query(...) 会返回在 arr[left], arr[left+1], ..., arr[right] 中至少出现阈值次数 threshold 的元素，如果不存在这样的元素，就返回 -1。

  提示：
    1 <= arr.length <= 20000
    1 <= arr[i] <= 20000
    对于每次查询，0 <= left <= right < len(arr)
    对于每次查询，2 * threshold > right - left + 1
    查询次数最多为 10000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/online-majority-element-in-subarray
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type MajorityChecker struct {
	a []int         // source array
	m map[int][]int // record num's index set
	t []*entity     // segment tree
}

func Constructor(arr []int) MajorityChecker {
	mc := MajorityChecker{}
	mc.a = arr
	mc.m = make(map[int][]int)
	for i := 0; i < len(arr); i++ {
		mc.m[arr[i]] = append(mc.m[arr[i]], i)
	}
	n := len(arr)
	mc.t = make([]*entity, calTreeSize(n))
	buildTree(arr, mc.t, 0, 0, n-1)
	return mc
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
	res := queryTree(this.t, 0, 0, len(this.a)-1, left, right)
	if res == nil || res.value == 0 {
		return -1
	}
	l := sort.Search(len(this.m[res.value]), func(i int) bool { return left <= this.m[res.value][i] })
	r := sort.Search(len(this.m[res.value]), func(i int) bool { return right < this.m[res.value][i] })
	if r-l < threshold {
		return -1
	}
	return res.value
}

func buildTree(arr []int, t []*entity, index, l, r int) {
	if l == r {
		t[index] = &entity{value: arr[l], count: 1}
		return
	}
	m := l + (r-l)>>1
	childL := index<<1 + 1
	childR := childL + 1
	buildTree(arr, t, childL, l, m)
	buildTree(arr, t, childR, m+1, r)
	t[index] = addSegTree(t[childL], t[childR])
}

func queryTree(t []*entity, index, l, r, L, R int) *entity {
	if l > R || r < L {
		return nil
	} else if L <= l && r <= R {
		return t[index]
	}
	m := l + (r-l)>>1
	childL := index<<1 + 1
	childR := childL + 1
	return addSegTree(queryTree(t, childL, l, m, L, R), queryTree(t, childR, m+1, r, L, R))
}

type entity struct {
	value int
	count int
}

func calTreeSize(arrLen int) (size int) {
	for i := 1; i>>1 < arrLen; i <<= 1 {
		size += i
	}
	return
}

func addSegTree(l, r *entity) *entity {
	if l == nil && r == nil {
		return nil
	} else if l == nil {
		return r
	} else if r == nil {
		return l
	}
	var val, cnt int
	if l.value == r.value {
		val, cnt = l.value, l.count+r.count
	} else if l.count < r.count {
		val, cnt = r.value, r.count-l.count
	} else {
		val, cnt = l.value, l.count-r.count
	}
	return &entity{value: val, count: cnt}
}

/**
 * Your MajorityChecker object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,threshold);
 */
