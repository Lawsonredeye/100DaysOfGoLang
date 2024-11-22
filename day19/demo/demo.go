package split

import "strings"

// Similar to the split func in the strings library.
func Split(s, sep string) []string {
	var result []string

	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	return append(result, s)
}
