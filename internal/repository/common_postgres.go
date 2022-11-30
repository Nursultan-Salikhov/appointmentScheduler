package repository

import "strings"

func correctDateFormat(date string) string {
	before, _, found := strings.Cut(date, "T")
	if found {
		return before
	}
	return date
}

func correctTimeFormat(time string) string {
	_, after, found := strings.Cut(time, "T")
	if found {
		after = strings.TrimRight(after, "00Z")
		return strings.TrimRight(after, ":")
	}
	return time
}
