package utils

import "unicode"

func HasSpecialChars(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			continue
		}
		return true
	}
	return false
}

func ToTitle(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(s)
	for i, r := range runes {
		if i == 0 || unicode.IsSpace(runes[i-1]) {
			runes[i] = unicode.ToTitle(r)
		}
	}
	return string(runes)
}
