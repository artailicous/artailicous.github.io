// Package weather provides a simple utility for tracking and reporting the current
// weather condition and location.
//
// The Forecast function updates the package's exported state variables.
package weather

var (
    // CurrentCondition holds the textual description of the last reported weather (e.g., "Sunny").
	CurrentCondition string
    // CurrentLocation holds the name of the city for the last reported weather (e.g., "Stockholm").
	CurrentLocation  string
)
// Forecast sets the current city and weather condition in the package's global state
// and returns a formatted string summarizing the new forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
