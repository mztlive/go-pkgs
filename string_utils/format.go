package string_utils

import "unicode"

// toSnakeCase converts a string to snake_case.
func ToSnakeCase(s string) string {
	var result []rune
	for i, v := range s {
		if unicode.IsUpper(v) {
			// Check if this uppercase letter is the start of a new word.
			// This happens if the previous letter is not uppercase, or if the next letter is lowercase.
			if i != 0 && (!unicode.IsUpper(rune(s[i-1])) || (i+1 < len(s) && unicode.IsLower(rune(s[i+1])))) {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(v))
		} else {
			result = append(result, v)
		}
	}
	return string(result)
}
