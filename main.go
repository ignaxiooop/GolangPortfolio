package main

import (
	"fmt"  // Import the fmt package for formatted I/O operations
	"time" // Import the time package to handle time-related functionality
)

func main() {
	// Get the current time
	ahora := time.Now()

	// Format the time to display hours, minutes, AM/PM, and timezone
	// "15:04 PM MST" specifies:
	// "15" for hour in 24-hour format, which gets converted to 12-hour format in the output
	// "04" for minutes
	// "PM" automatically displays AM or PM
	// "MST" displays the time zone abbreviation
	horaformateada := ahora.Format("15:04 PM MST")

	fmt.Println("Hey! The current time is: " + horaformateada)

}
