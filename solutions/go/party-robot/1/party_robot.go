package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	// Uses fmt.Sprintf to create the welcome message.
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	// Uses fmt.Sprintf to create the birthday message.
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	// The first line is the welcome message, reusing the Welcome function.
	welcomeLine := Welcome(name)

	// The table number needs to be 3 digits with leading zeroes if necessary.
	// We use the format specifier "%03d" for this.
	tableLine := fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.", table, direction, distance)

	// The distance needs to be limited to 1 digit precision.
	// We use the format specifier "%.1f" for this.

	// The last line mentions the seatmate.
	neighborLine := fmt.Sprintf("You will be sitting next to %s.", neighbor)

	// Combine all three lines with newline characters "\n".
	return fmt.Sprintf("%s\n%s\n%s", welcomeLine, tableLine, neighborLine)
}