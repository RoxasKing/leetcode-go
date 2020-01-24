package sort

func ShellSort(array []int) {
	gap := len(array) / 2
	var i, j, tmp int
	for gap > 0 {
		for i = gap; i < len(array); i++ {
			tmp = array[i]
			for j = i - gap; j >= 0 && array[j] > tmp; j = j - gap {
				array[j+gap] = array[j]
			}
			array[j+gap] = tmp
		}
		gap /= 2
	}
}
