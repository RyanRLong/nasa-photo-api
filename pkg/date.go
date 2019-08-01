package nasaphotoapi

import (
	"fmt"
	"strconv"
	"time"
)

// Date represents day, month, and year
type Date struct {
	Year  string
	Month string
	Day   string
}

func (date Date) String() string {
	return date.Year + "-" + date.Month + "-" + date.Day
}

// IsValid returns true if the struct contains a valid day, month, and year
func (date Date) IsValid() bool {
	const shortForm = "20060102"

	day := padToTwoDigits(date.Day)
	month := padToTwoDigits(date.Month)
	year := date.Year

	_, err := time.Parse(shortForm, year+month+day)
	if err != nil {
		return false
	}
	return true
}

func padToTwoDigits(value string) string {
	intValue, _ := strconv.Atoi(value)
	return fmt.Sprintf("%02d", intValue)
}

// DateError is an error implementation that holds the offending string
// with a message
type DateError struct {
	DateString Date
	Message    string
}

func (e DateError) Error() string {
	return fmt.Sprintf("%v: %v", e.Message, e.DateString)
}
