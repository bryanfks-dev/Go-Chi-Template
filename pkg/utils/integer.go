package utils

import (
	"strconv"
	"strings"
)

func ParseIntPtr(s string) *int {
	intg, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &intg
}

func ParseIntOrDefault(value string, defaultValue int) int {
	if value == "" {
		return 0
	}
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return parsedValue
}

func ParseCommaSeparatedInts(s string) []int {
	var ints []int
	for part := range strings.SplitSeq(s, ",") {
		if intg, err := strconv.Atoi(part); err == nil {
			ints = append(ints, intg)
		}
	}
	return ints
}
