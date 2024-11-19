package main

import(
	"time"
	"fmt"	
)

type formattedDateTime struct{
	Date time.Time
	Time time.Time
}

func timeFormatter(inputDate string, inputTime string) (formattedDateTime,error){
	const dateFormat = "2006-01-02"
	const timeFormat = "15:04:05"
	
	// Parse date and time
	datePart, err := time.Parse(dateFormat, inputDate)
	if err != nil {
		return formattedDateTime{}, fmt.Errorf("invalid date format: %v", err)
	}

	timePart, err := time.Parse(timeFormat, inputTime)
	if err != nil {
		return formattedDateTime{}, fmt.Errorf("invalid time format: %v", err)
	}
	return formattedDateTime{
		Date: datePart,
		Time: timePart,
	},nil
}
