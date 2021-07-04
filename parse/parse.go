package parse

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Define minimum and maximum allowed timestamps as
// specified by pdf [Coding_test.pdf].
// Valid date values are equal to or above MinDate and
// equal to or below MaxDate
var (
	MinDate = time.Date(
		1899, time.December, 31, 0, 0, 0, 0, time.UTC)

	MaxDate = time.Date(3000, time.January, 1, 0, 0, 0, 0, time.UTC)
)

// ParseTime Validates string and returns time.Time
func ParseTime(input string) (time.Time, error) {
	// validate input with regex
	match, err := regexp.MatchString(`^[0-9]{1,2}\/[0-9]{1,2}\/[0-9]{1,4}$`, input)
	if err != nil {
		return time.Time{}, err
	}
	if !match {
		return time.Time{}, errors.New("invalid time format")
		// return error
	}
	inputArray := strings.Split(input, "/")
	for location, dateStamp := range inputArray {
		// lazy validation of string
		// if less 2 strings per position format from 1 to 01
		if len(dateStamp) < 2 {
			tempFmt := fmt.Sprintf("0%s", dateStamp)
			// modify inputArray with the new format
			inputArray[location] = tempFmt

		}

	}

	newFormattedInput := strings.Join(inputArray, "/")
	parsedDate, err := time.Parse("02/01/2006", newFormattedInput)
	if err != nil {
		return time.Time{}, err
	}
	// validate date range
	if !parsedDate.After(MinDate) && parsedDate.Before(MaxDate) {

		return time.Time{}, errors.New("error dates outside of acceptable range")
	}
	return parsedDate, nil
}

// CalculateDayDifference Calculates the days difference between input 1 and input 2
func CalculateDayDifference(input1 time.Time, input2 time.Time) (int, error) {
	// input1 and input 2 are the sameday
	if input1.Equal(input2) {
		return 0, nil

		// input1 is after input2
	} else if input1.After(input2) {
		days := input1.Sub(input2).Hours() / 24
		return int(days) + 1, nil

		// input1 is before input2
	} else if input1.Before(input2) {
		days := input2.Sub(input1).Hours() / 24
		return int(days) - 1, nil
	} else {
		return 0, errors.New("error unhandled condition")
	}
}
