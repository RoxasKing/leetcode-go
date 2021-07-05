package main

// Tags:
// Hash + Graph
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	dict := make(map[string][]endPoint)
	for i, equation := range equations {
		dict[equation[0]] = append(dict[equation[0]], endPoint{name: equation[1], value: values[i]})
		dict[equation[1]] = append(dict[equation[1]], endPoint{name: equation[0], value: 1 / values[i]})
	}
	out := make([]float64, len(queries))
	for i, query := range queries {
		out[i] = search(dict, map[string]bool{}, query, 1)
	}
	return out
}

type endPoint struct {
	name  string
	value float64
}

func search(dict map[string][]endPoint, walked map[string]bool, query []string, temp float64) float64 {
	walked[query[0]] = true
	for _, endPoint := range dict[query[0]] {
		if endPoint.name == query[1] {
			return temp * endPoint.value
		}
		if walked[endPoint.name] {
			continue
		}
		if ret := search(dict, walked, []string{endPoint.name, query[1]}, temp*endPoint.value); ret != -1.0 {
			return ret
		}
	}
	walked[query[0]] = false
	return -1.0
}
