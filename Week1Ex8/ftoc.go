// Package ftoc converts temperature given in deg Fahrenheit to equivalent
// value in deg Celsius.
package ftoc

// Convertftoc takes input as temperate in deg Fahrenheit and returns the temperature
// converted to deg Celsius
func Convertftoc(tempf float64) float64 {

	return ((tempf - 32) * 5 / 9)
}
