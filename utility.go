package main

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

func isWhitespaceOrBlank(str string) bool {
	return strings.TrimSpace(str) == ""
}
func isValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func parseDateWithFormats(dateString string) (time.Time, error) {
	formats := []string{
		time.RFC3339,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC822,
		time.RFC822Z,
		time.RFC3339Nano,
		"2006-01-02T15:04:05-07:00",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05",
		"02 Jan 2006 15:04:05 MST",
		"02 Jan 2006 15:04:05 -0700",
		"02 Jan 2006 15:04:05",
		"02 Jan 2006 15:04 MST",
		"02 Jan 2006 15:04 -0700",
		"02 Jan 2006 15:04",
		"02 Jan 2006 MST",
		"02 Jan 2006 -0700",
		"02 Jan 2006",
		"Mon, 02 Jan 2006 15:04:05 -0700",
		"Mon, 02 Jan 2006 15:04:05 +0000", // Additional layout
		"2006-01-02T15:04:05+00:00",       // New layout for "2021-03-25T14:00:00+00:00"
	}

	// Remove leading and trailing whitespace characters
	dateString = strings.TrimSpace(dateString)

	var parsedTime time.Time
	var err error
	for _, layout := range formats {
		parsedTime, err = time.Parse(layout, dateString)
		if err == nil {
			return parsedTime, nil // Parsing succeeded, return the parsed time
		}
	}

	// If parsing fails for all formats, return the last encountered error
	return time.Time{}, fmt.Errorf("failed to parse date: %v", err)
}
