package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ewhal/time-distance-golang/parse"
)

func main() {

	// parsed input is put here
	dates := make([]time.Time, 0)

	// new scanner taking input from stdin
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// print input text
		fmt.Print("Enter date in format 01/01/2001 or 1/1/2001: ")

		if !scanner.Scan() {
			break
		}

		// get text from stdin
		text := scanner.Text()

		// if not empty string call ParseTime
		if len(text) != 0 {

			// trim input string
			trimmed := strings.TrimSpace(text)
			timeInput, err := parse.ParseTime(trimmed)
			if err != nil {
				fmt.Println(err)
				break
			}
			dates = append(dates, timeInput)
		}
		// once we have 2 dates call parse time
		if len(dates) == 2 {
			dateDifference, err := parse.CalculateDayDifference(dates[0], dates[1])
			if err != nil {
				fmt.Println(err)
				break

			}
			fmt.Printf("The amount of days between %s and %s is %d \n", dates[0].String(), dates[1].String(), dateDifference)
			break
		}

		if err := scanner.Err(); err != nil {
			break
		}

	}

}
