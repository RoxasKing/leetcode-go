package main

// Tags:
// Hash

func numUniqueEmails(emails []string) int {
	set := map[string]struct{}{}
	for _, email := range emails {
		arr := make([]byte, 0, len(email))
		i := 0
		for ; email[i] != '+' && email[i] != '@'; i++ {
			if email[i] != '.' {
				arr = append(arr, email[i])
			}
		}
		for ; email[i] != '@'; i++ {
		}
		arr = append(arr, '@')
		for i++; i < len(email); i++ {
			arr = append(arr, email[i])
		}
		set[string(arr)] = struct{}{}
	}
	return len(set)
}
