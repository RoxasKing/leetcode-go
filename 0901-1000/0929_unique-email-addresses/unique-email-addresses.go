package main

// Difficulty:
// Easy

// Tags:
// Hash

func numUniqueEmails(emails []string) int {
	set := map[string]struct{}{}
	for _, email := range emails {
		a := []byte(email)
		i, j := 0, 0
		for ; email[j] != '+' && email[j] != '@'; j++ {
			if email[j] != '.' {
				a[i] = email[j]
				i++
			}
		}
		for ; email[j] != '@'; j++ {
		}
		for ; j < len(email); j++ {
			a[i] = email[j]
			i++
		}
		set[string(a[:i])] = struct{}{}
	}
	return len(set)
}
