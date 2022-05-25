package main

import (
	"errors"
	"flag"
	"fmt"
	"time"
)

var (
	ErrNoTimeOrDate            = errors.New("No time or datetime given")
	ErrClashingDateAndDatetime = errors.New("Both time and datetime given")

	dateFormat    string
	inputTime     string
	inputDatetime string
	inputLocation string
)

const (
	InputDatetimeFormat = "2006-01-02T15:04"
	InputTimeFormat     = "15:04"
)

func main() {
	t, err := parseArgs()
	if err != nil {
		usage()
		return
	}

	fmt.Printf("Reference Date:\t%s\n", t.Format(dateFormat))

	loc, _ := time.LoadLocation("America/Los_Angeles")
	fmt.Printf("West Coast:\t%s\n", t.In(loc).Format(dateFormat))

	loc, _ = time.LoadLocation("America/New_York")
	fmt.Printf("East Coast:\t%s\n", t.In(loc).Format(dateFormat))

	loc, _ = time.LoadLocation("UTC")
	fmt.Printf("UTC:\t\t%s\n", t.In(loc).Format(dateFormat))

	loc, _ = time.LoadLocation("Europe/Madrid")
	fmt.Printf("Spain:\t\t%s\n", t.In(loc).Format(dateFormat))

	loc, _ = time.LoadLocation("Asia/Kolkata")
	fmt.Printf("India:\t\t%s\n", t.In(loc).Format(dateFormat))

}

func parseArgs() (*time.Time, error) {

	flag.StringVar(&dateFormat, "f", time.RFC3339, "A golang style time format used to format the output")
	flag.StringVar(&inputTime, "t", "", fmt.Sprintf("A 24-Hour formatted time like %s", InputTimeFormat))
	flag.StringVar(&inputDatetime, "d", "", fmt.Sprintf("An ISO 8601 Datime time like %s", InputDatetimeFormat))
	flag.StringVar(&inputLocation, "l", time.Local.String(), "An IANA Time Zone like America/New_York")

	flag.Parse()
	now := time.Now()
	if inputDatetime == "" && inputTime == "" {

		return &now, nil

	}
	if inputDatetime != "" && inputTime != "" {
		return nil, ErrClashingDateAndDatetime
	}

	loc, err := time.LoadLocation(inputLocation)
	if err != nil {
		return nil, err
	}

	var parsedTime time.Time
	if inputDatetime != "" {
		parsedTime, err = time.Parse(InputDatetimeFormat, inputDatetime)
		if err != nil {
			return nil, err
		}
		parsedTime = time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, loc)
	} else {
		parsedTime, err = time.Parse(InputTimeFormat, inputTime)
		if err != nil {
			return nil, err
		}
		parsedTime = time.Date(now.Year(), now.Month(), now.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, loc)
	}

	return &parsedTime, nil
}

func usage() {
	fmt.Println("Usage: tz -t 18:00")
	flag.PrintDefaults()
}
