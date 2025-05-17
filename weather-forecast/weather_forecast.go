// Package weather provides tools to convert.
// temperatures to and from Kelvin.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current Location.
var CurrentLocation string

// Forecast makes a prediction for a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
