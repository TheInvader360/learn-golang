package genconv

// CelsiusToFahrenheit converts a Celsius temperature to Fahrenheit.
func CelsiusToFahrenheit(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FahrenheitToCelsius converts a Fahrenheit temperature to Celsius.
func FahrenheitToCelsius(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CelsiusToKelvin converts a Celsius temperature to Kelvin.
func CelsiusToKelvin(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FootToMeter converts a length from feet to meters
func FootToMeter(f Foot) Meter { return Meter(f * 0.3048) }

// MeterToFoot converts a length from meters to feet
func MeterToFoot(m Meter) Foot { return Foot(m * 3.28084) }

// PoundToKilogram converts a mass from pounds to kilograms
func PoundToKilogram(p Pound) Kilogram { return Kilogram(p * 0.453592) }

// KilogramToPound converts a mass from kilograms to pounds
func KilogramToPound(k Kilogram) Pound { return Pound(k * 2.20462) }
