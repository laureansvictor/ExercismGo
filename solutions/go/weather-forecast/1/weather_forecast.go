// Package weather containing the weather information of a city.
package weather

var (
	// CurrentCondition a variable containing the current condition.
	CurrentCondition string
	// CurrentLocation a variable containing the location.
	CurrentLocation string
)

// Forecast is a function that returns the forecast of a city with its condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
