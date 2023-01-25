// Package weather. Provides the current
// state of zones.
package weather

// CurrentCondition contains the current condition of the time.
var CurrentCondition string
// CurrentLocation contains the current location.
var CurrentLocation string

// Forecast function
// Due to the current location and the current city, returns some info.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
