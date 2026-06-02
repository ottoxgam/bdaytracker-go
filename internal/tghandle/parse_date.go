package tghandle

import (
	"fmt"
	"strings"
	"time"
)

const numberOfMonths = 12

var daysBefore = [...]int{
	0,
	31,
	31 + 28,
	31 + 28 + 31,
	31 + 28 + 31 + 30,
	31 + 28 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30,
	31 + 28 + 31 + 30 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30 + 31,
}

type dateParseError int

const (
	dateParseOK dateParseError = iota
	dateParseInvalid
	dateParseWrongDays
)

// parseBirthday parses a birthday from "MM.DD" or "MM/DD" format.
func parseBirthday(s string) (month, day int, result dateParseError) {
	normalized := strings.ReplaceAll(s, "/", ".")
	_, err := fmt.Sscanf(normalized, "%d.%d", &month, &day)
	if err != nil || month < 1 || month > numberOfMonths {
		return 0, 0, dateParseInvalid
	}

	daysInMonth := daysBefore[month] - daysBefore[month-1]
	if month == int(time.February) {
		daysInMonth++
	}
	if day < 1 || day > daysInMonth {
		return 0, 0, dateParseWrongDays
	}

	return month, day, dateParseOK
}
