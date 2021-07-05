package main

func validateStackSequences(pushed []int, popped []int) bool {
	help := make([]int, 0, len(pushed))
	for len(popped) > 0 {
		for len(pushed) > 0 && (len(help) == 0 || help[len(help)-1] != popped[0]) {
			help = append(help, pushed[0])
			pushed = pushed[1:]
		}
		if help[len(help)-1] != popped[0] {
			return false
		}
		help = help[:len(help)-1]
		popped = popped[1:]
	}
	return true
}
