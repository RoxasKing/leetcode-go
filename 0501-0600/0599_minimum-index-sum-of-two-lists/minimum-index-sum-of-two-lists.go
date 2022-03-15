package main

// Difficulty:
// Easy

// Tags:
// Hash

func findRestaurant(list1 []string, list2 []string) []string {
	set := map[string]int{}
	for i, name := range list1 {
		set[name] = i
	}
	out, sum := []string{}, 1<<31-1
	for i, name := range list2 {
		if j, ok := set[name]; !ok {
			continue
		} else if sum == i+j {
			out = append(out, name)
		} else if sum > i+j {
			sum = i + j
			out = []string{name}
		}
	}
	return out
}
