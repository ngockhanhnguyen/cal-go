package main

import (
	"fmt"
	"time"
)

func main() {
	printMonthAndYear()
	printDaysOfWeek()
	printDaysOfMonth()
}

func printMonthAndYear() {
	currentTime := time.Now()
	month := currentTime.Month().String()
	year := currentTime.Year()
	stringRepresentation := fmt.Sprintf("%s %d", month, year)

	// Print the result
	fmt.Println(stringRepresentation)
}

func printDaysOfWeek() {
}

func printDaysOfMonth() {
}
