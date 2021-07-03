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

	dates := make([]time.Time, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter date in format 01/01/2001 or 1/1/2001: ")

		scanner.Scan()

		text := scanner.Text()

		if len(text) != 0 {

			trimmed := strings.Trim(text, "")
			timeInput, err := parse.ParseTime(trimmed)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(timeInput.String())
			dates = append(dates, timeInput)
		}
		if len(dates) == 2 {
			dateDifferent, err := parse.CalculateDayDifference(dates[0], dates[1])
			if err != nil {
				fmt.Println(err)
				break

			}
			fmt.Println(dateDifferent)
			break
		}

	}

}
