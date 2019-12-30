package sort

import (
	"fmt"
	"testing"
)

var nums = []int{7, 10, 8, 9, 3, 2, 1, 11, 6, 5, 4, 12, 14, 13, 0}

func TestBubbleSort(t *testing.T) {
	array := make([]int, len(nums))
	copy(array, nums)
	fmt.Printf("Before bubble sort : %v\n", array)
	BubbleSort(array)
	fmt.Printf("After bubble sort  : %v\n", array)
}

func TestSelectionSort(t *testing.T) {
	array := make([]int, len(nums))
	copy(array, nums)
	fmt.Printf("Before selection sort : %v\n", array)
	SelectionSort(array)
	fmt.Printf("After selection sort  : %v\n", array)
}

func TestInsertionSort(t *testing.T) {
	array := make([]int, len(nums))
	copy(array, nums)
	fmt.Printf("Before insertion sort : %v\n", array)
	InsertionSort(array)
	fmt.Printf("After insertion sort  : %v\n", array)
}

func TestShellSort(t *testing.T) {
	array := make([]int, len(nums))
	copy(array, nums)
	fmt.Printf("Before shell sort : %v\n", array)
	ShellSort(array)
	fmt.Printf("After shell sort  : %v\n", array)
}

func TestMergeSort(t *testing.T) {
	array := make([]int, len(nums))
	copy(array, nums)
	fmt.Printf("Before merge sort : %v\n", array)
	MergeSort(array)
	fmt.Printf("After merge sort  : %v\n", array)
}

func TestQuickSort(t *testing.T) {
	array := make([]int, len(nums))
	copy(array, nums)
	fmt.Printf("Before quick sort : %v\n", array)
	QuickSort(array)
	fmt.Printf("After quick sort  : %v\n", array)
}

func TestHeapSort(t *testing.T) {
	array := make([]int, len(nums))
	copy(array, nums)
	fmt.Printf("Before heap sort : %v\n", array)
	HeapSort(array)
	fmt.Printf("After heap sort  : %v\n", array)
}

func TestCountingSort(t *testing.T) {
	array := make([]int, len(nums))
	copy(array, nums)
	fmt.Printf("Before counting sort : %v\n", array)
	CountingSort(array)
	fmt.Printf("After counting sort  : %v\n", array)
}

func TestBucketSort(t *testing.T) {
	array := make([]int, len(nums))
	copy(array, nums)
	fmt.Printf("Before bucket sort : %v\n", array)
	BucketSort(array)
	fmt.Printf("After bucket sort  : %v\n", array)
}

func TestRadixSort(t *testing.T) {
	array := make([]int, len(nums))
	copy(array, nums)
	fmt.Printf("Before radix sort : %v\n", array)
	RadixSort(array)
	fmt.Printf("After radix sort  : %v\n", array)
}
