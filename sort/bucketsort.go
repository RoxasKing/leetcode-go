package sort

func BucketSort(array []int) {
	if len(array) == 0 {
		return
	}
	min, max := array[0], array[0]
	for i := range array {
		if array[i] < min {
			min = array[i]
		} else if array[i] > max {
			max = array[i]
		}
	}
	factor := 1
	buckets := make([][]int, (max-min)/factor+1)
	for i := range array {
		index := (array[i] - min) / factor
		buckets[index] = append(buckets[index], array[i])
	}
	var index int
	for i := range buckets {
		if len(buckets[i]) <= 0 {
			continue
		}
		InsertionSort(buckets[i])
		for j := range buckets[i] {
			array[index] = buckets[i][j]
			index++
		}
	}
}
