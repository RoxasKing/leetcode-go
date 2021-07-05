package main

// Tags:
// DFS
func verifyPostorder(postorder []int) bool {
	return isValid(postorder, 0, len(postorder)-1)
}

func isValid(postorder []int, l, r int) bool {
	if l >= r {
		return true
	}
	rootVal := postorder[r]
	i := l
	for i < r && postorder[i] <= rootVal {
		i++
	}
	for j := i; j < r; j++ {
		if postorder[j] < rootVal {
			return false
		}
	}
	return isValid(postorder, l, i-1) && isValid(postorder, i, r-1)
}
