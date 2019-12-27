package My_LeetCode_In_Go

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength&1 == 1 {
		return float64(findKthElem(nums1, nums2, totalLength/2+1))
	}
	return 0.5 * float64(
		findKthElem(nums1, nums2, totalLength/2)+
			findKthElem(nums1, nums2, totalLength/2+1))

}

//1<=k<=len(nums1)+len(nums2)
func findKthElem(nums1, nums2 []int, k int) int {
	var i, j, l1 int
	for {
		if len(nums1) > len(nums2) {
			nums1, nums2 = nums2, nums1
		}
		l1 = len(nums1)
		if l1 == 0 {
			return nums2[k-1]
		}
		if k == 1 {
			return min(nums1[0], nums2[0])
		}
		if k/2 > l1 {
			i = l1
		} else {
			i = k / 2
		}
		j = k - i
		// i,j <= k/2 <= l1 <= l2
		if nums1[i-1] > nums2[j-1] {
			nums2 = nums2[j:]
			k -= j
		} else if nums2[j-1] > nums1[i-1] {
			nums1 = nums1[i:]
			k -= i
		} else {
			return nums1[i-1]
		}
	}
}
