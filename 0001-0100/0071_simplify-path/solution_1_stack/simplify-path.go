package main

// Difficulty:
// Medium

// Tags:
// Stack

func simplifyPath(path string) string {
	stack := []string{}
	n := len(path)
	for i := 0; i < n; {
		for ; i < n && path[i] == '/'; i++ {
		}
		if i < n && path[i] == '.' && (i+1 == n || path[i+1] == '/') {
			i++
		} else if i < n && path[i] == '.' && (i+1 < n && path[i+1] == '.') && (i+2 == n || path[i+2] == '/') {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			i += 2
		} else if i < n {
			j := i
			for ; j < n && path[j] != '/'; j++ {
			}
			stack = append(stack, path[i:j])
			i = j
		}
	}
	var out string
	for _, name := range stack {
		out += "/" + name
	}
	if len(out) == 0 {
		return "/"
	}
	return out
}
