package main

func memLeak(memory1 int, memory2 int) []int {
	for i := 1; ; i++ {
		if memory1 >= memory2 && memory1 >= i {
			memory1 -= i
		} else if memory1 < memory2 && memory2 >= i {
			memory2 -= i
		} else {
			return []int{i, memory1, memory2}
		}
	}
}
