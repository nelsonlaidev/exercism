package twofer

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	result := "you"

	if name != "" {
		result = name
	}

	return fmt.Sprintf("One for %s, one for me.", result)
}
