package utils

import (
	"regexp"
	"strings"
	"unicode"
)

func HasSpecialChars(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			continue
		}
		return true
	}
	return false
}

func ToLowerKebab(s string) string {
	if s == "" {
		return s
	}

	re := regexp.MustCompile(`\s+`)
	normalized := re.ReplaceAllString(s, " ")
	normalized = strings.TrimSpace(normalized)
	normalized = strings.ReplaceAll(normalized, " ", "-")
	normalized = strings.ToLower(normalized)
	return normalized
}

func Normalize(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(s)
	for i, r := range runes {
		if i == 0 || unicode.IsSpace(runes[i-1]) {
			runes[i] = unicode.ToTitle(r)
			continue
		}

		runes[i] = unicode.ToLower(r)
	}

	re := regexp.MustCompile(`\s+`)
	normalized := re.ReplaceAllString(string(runes), " ")
	normalized = strings.TrimSpace(normalized)
	return normalized
}

func ToTitle(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(s)
	for i, r := range runes {
		if i == 0 || unicode.IsSpace(runes[i-1]) {
			runes[i] = unicode.ToTitle(r)
			continue
		}

		runes[i] = unicode.ToLower(r)
	}
	return strings.ReplaceAll(string(runes), " ", "")
}

func ToCamel(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(s)
	for i, r := range runes {
		if i == 0 || !unicode.IsSpace(runes[i-1]) {
			runes[i] = unicode.ToLower(r)
			continue
		}

		runes[i] = unicode.ToTitle(r)
	}
	return string(runes)
}
