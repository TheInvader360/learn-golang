package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/TheInvader360/learn-golang/gopl-book/ch2/genconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		fahrenheit := genconv.Fahrenheit(t)
		celsius := genconv.Celsius(t)
		foot := genconv.Foot(t)
		meter := genconv.Meter(t)
		pound := genconv.Pound(t)
		kilogram := genconv.Kilogram(t)

		fmt.Printf("%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n",
			fahrenheit, genconv.FahrenheitToCelsius(fahrenheit),
			celsius, genconv.CelsiusToFahrenheit(celsius),
			celsius, genconv.CelsiusToKelvin(celsius),
			foot, genconv.FootToMeter(foot),
			meter, genconv.MeterToFoot(meter),
			pound, genconv.PoundToKilogram(pound),
			kilogram, genconv.KilogramToPound(kilogram))
	}
}
