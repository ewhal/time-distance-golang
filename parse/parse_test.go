package parse

import (
	"fmt"
	"testing"
	"time"
)

// DateStruct
type DateStruct struct {
	// Input1
	Input1 string
	// Input2
	Input2 string
	// ExpectedResultParse1
	ExpectedResultParse1 time.Time
	// ExpectedResultParse2
	ExpectedResultParse2 time.Time
	// ExpectedResultDayDifference
	ExpectedResultDayDifference int
	//ShouldError
	ShouldError bool
}

// TestDates
var TestDates []DateStruct = []DateStruct{
	{
		Input1: "01/01/2001",
		Input2: "03/01/2001",
		ExpectedResultParse1: time.Date(
			2001, time.January, 01, 0, 0, 0, 0, time.UTC),
		ExpectedResultParse2: time.Date(
			2001, time.January, 03, 0, 0, 0, 0, time.UTC),
		ExpectedResultDayDifference: 1,
		ShouldError:                 false,
	},
	{
		Input1: "01/01/1900",
		Input2: "31/12/2999",
		ExpectedResultParse1: time.Date(
			1900, time.January, 01, 0, 0, 0, 0, time.UTC),
		ExpectedResultParse2: time.Date(
			2999, time.December, 31, 0, 0, 0, 0, time.UTC),

		ExpectedResultDayDifference: 106750,
		ShouldError:                 false,
	},
	{
		Input1: "01/01/1899",
		Input2: "31/12/3000",
		ExpectedResultParse1: time.Date(
			0001, time.January, 01, 0, 0, 0, 0, time.UTC),
		ExpectedResultParse2: time.Date(
			3000, time.December, 31, 0, 0, 0, 0, time.UTC),

		ExpectedResultDayDifference: 0,
		ShouldError:                 true,
	},
	{
		Input1: "2/6/1983",
		Input2: "22/6/1983",
		ExpectedResultParse1: time.Date(
			1983, time.June, 02, 0, 0, 0, 0, time.UTC),
		ExpectedResultParse2: time.Date(
			1983, time.June, 22, 0, 0, 0, 0, time.UTC),

		ExpectedResultDayDifference: 19,
		ShouldError:                 false,
	},
	{
		Input1: "4/7/1984",
		Input2: "25/12/1984",
		ExpectedResultParse1: time.Date(
			1984, time.July, 04, 0, 0, 0, 0, time.UTC),
		ExpectedResultParse2: time.Date(
			1984, time.December, 25, 0, 0, 0, 0, time.UTC),

		ExpectedResultDayDifference: 173,
		ShouldError:                 false,
	},
	// Example given by pdf has incorrect expected day difference
	{
		Input1: "3/1/1989",
		Input2: "3/8/1983",
		ExpectedResultParse1: time.Date(
			1989, time.January, 03, 0, 0, 0, 0, time.UTC),
		ExpectedResultParse2: time.Date(
			1983, time.August, 03, 0, 0, 0, 0, time.UTC),

		ExpectedResultDayDifference: 2036,
		ShouldError:                 true,
	},
	// Correct final day difference test
	{
		Input1: "3/1/1989",
		Input2: "3/8/1983",
		ExpectedResultParse1: time.Date(
			1989, time.January, 03, 0, 0, 0, 0, time.UTC),
		ExpectedResultParse2: time.Date(
			1983, time.August, 03, 0, 0, 0, 0, time.UTC),

		ExpectedResultDayDifference: 1981,
		ShouldError:                 false,
	},
}

var BrokenTimeFormats = []string{
	"May 8, 2009 5:57:51 PM",
	"2014/04/08 22:05",
	"2014/04",
	"08/04",
	"2006-01-02T15:04:05+0000",
	"2013-04-01 22:43",
	"2014",
	"08.21.71",
	"1384216367189",
}

// TestTimeFormat tests various invalid time inputs for this project
func TestTimeFormat(test *testing.T) {
	for _, tt := range BrokenTimeFormats {
		testname := fmt.Sprintf("parse %s", tt)
		test.Run(testname, func(t *testing.T) {
			result, err := ParseTime(tt)
			if err == nil {
				t.Errorf("error expected error but got %s", result.String())
			}
		})
	}
}

// TestParseTime test all example time inputs given by the project specifications
func TestParseTime(test *testing.T) {
	for _, tt := range TestDates {
		testname := fmt.Sprintf("%s,%s", tt.Input1, tt.Input2)
		test.Run(testname, func(t *testing.T) {
			result, err := ParseTime(tt.Input1)
			if err != nil && !tt.ShouldError {
				t.Error(err)
			}

			if result != tt.ExpectedResultParse1 {
				t.Errorf("got %s, wanted %s", result.String(), tt.ExpectedResultParse1.String())
			}
			result2, err := ParseTime(tt.Input2)
			if err != nil && !tt.ShouldError {
				t.Error(err)
			}

			if result2 != tt.ExpectedResultParse2 {
				t.Errorf("got %s, wanted %s", result2.String(), tt.ExpectedResultParse2.String())
			}
		})
	}
}

// TestCalculateDayDifference caluclates the difference between the days
func TestCalculateDayDifference(test *testing.T) {
	for _, tt := range TestDates {
		testname := fmt.Sprintf("%s,%s", tt.ExpectedResultParse1.String(), tt.ExpectedResultParse1.String())
		test.Run(testname, func(t *testing.T) {
			result, err := CalculateDayDifference(tt.ExpectedResultParse1, tt.ExpectedResultParse2)
			if err != nil && !tt.ShouldError {
				t.Error(err)
			}

			if result != tt.ExpectedResultDayDifference && !tt.ShouldError {
				t.Errorf("got %d, wanted %d", result, tt.ExpectedResultDayDifference)
			}

		})
	}

}
