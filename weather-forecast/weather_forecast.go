// Package weather covers simple weather data - Forecast.
package weather

// CurrentCondition stores weather condition.
var CurrentCondition string  
// CurrentLocation Store Location of interest to obtain data.
var CurrentLocation string

// Forecast function returns weather condition for given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
