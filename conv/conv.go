package conv

var Fahrenheit float64
var Celsius float64
var Kelvin float64

// Konverterer Farhenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {
	Celsius = (Fahrenheit - 32) * (5 / 9)
	return Celsius
}

// Konverterer Celsius til Fahrenheit
func CelsiusToFahrenheit(value float64) float64 {
	Fahrenheit = Celsius*(9/5) + 32

	return Fahrenheit
}

// Konverterer Celsius til kelvin
func CelsiusToKelvin(value float64) float64 {
	Kelvin = Celsius + 273.15

	return Kelvin
}

// Konverterer Kelvin til celsius
func KelvinToCelsius(value float64) float64 {
	Celsius = Kelvin - 273.15

	return Celsius
}

// Konverterer Kelvin til Fahrenheit
func KelvinToFahrenheit(value float64) float64 {
	Fahrenheit = (Kelvin-273.15)*(9/5) + 32

	return Fahrenheit
}

// Konverterer Fahrenheit til kelvin
func FahrenheitToKelvin(value float64) float64 {
	Kelvin = (Fahrenheit-32)*(5/9) + 273.15

	return Kelvin
}
