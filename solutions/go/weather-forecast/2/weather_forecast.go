/* Package weather takes the latest data of two variables:
* condition
* city
to output combined string. */
package weather

var (
	// CurrentCondition is the current condition.
	CurrentCondition string
    // CurrentLocation is the city.
	CurrentLocation  string
)
// Forecast returns the latest data of the two variables above.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
