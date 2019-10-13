package Algorithm

import (
	"fmt"
	"testing"
)

func BubbleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	data := []int{5, 6, 7, 3, 8, 1, 9, 0, 3, 5}
	fmt.Println(data)
	BubbleSort(data)
	fmt.Println(data)
}
