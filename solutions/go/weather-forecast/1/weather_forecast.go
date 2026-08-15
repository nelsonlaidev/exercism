// Package weather provides weather-related functions.
package weather

var (
	// CurrentCondition describes the weather condition.
	CurrentCondition string
	// CurrentLocation indicates the weather location.
	CurrentLocation string
)

// Forecast forecasts the weather condition of a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
