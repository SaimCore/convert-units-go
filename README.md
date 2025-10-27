# convert-units-go ✨
convert-units-go is a Go package designed to simplify conversions between units of measurement, including weight, length, and temperature. It provides a user-friendly interface for developers to convert values ​​between different units.

[![GoDoc](https://godoc.org/github.com/SaimCore/convert-units-go?status.svg)](https://godoc.org/github.com/SaimCore/convert-units-go)
[![Go Report Card](https://goreportcard.com/badge/SaimCore/convert-units-go)](https://goreportcard.com/report/SaimCore/convert-units-go)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/SaimCore/convert-units-go/blob/main/LICENSE)


## Install 🛠️
```
go get github.com/SaimCore/convert-units-go
```
## How to use💡
```
package main

import (
	"fmt"

	"github.com/SaimCore/convert-units-go/convert/length"
	"github.com/SaimCore/convert-units-go/convert/temperature"
	"github.com/SaimCore/convert-units-go/convert/weight"
)

func main() {
	// --- Convert length units
	calc_length, _ := length.ConvertLength(1, "kilometer", "meter")
	fmt.Printf("length: %.3f\n", calc_length) // length: 1000.000

	m_to_Km, _ := length.MeterToKilometer(1000)
	fmt.Printf("%.3f Meter = %.3f Kilometer\n", 1000.0, m_to_Km) // 1000.000 Meter = 1.000 Kilometer

	// --- Convert weight units
	calc_weight, _ := weight.ConvertWeight(1, "kilogram", "gram")
	fmt.Printf("weight: %.3f\n", calc_weight) // weight: 1000.000

	kg_to_tonne, _ := weight.KilogramToTonne(1000)
	fmt.Printf("%.3f kilogram = %.3f tonne\n", 1000.0, kg_to_tonne) // 1000.000 kilogram = 1.000 tonne

	// --- Convert temperature units
	calc_temp, _ := temperature.ConvertTemperature(0, "celsius", "fahrenheit")
	fmt.Printf("temp: %.3f\n", calc_temp) // temp: 32.000

	c_to_k, _ := temperature.CelsiusToKelvin(1)
	fmt.Printf("%.3f celsius = %.3f kelvin\n", 1.0, c_to_k) // 1.000 celsius = 274.150 kelvin
}
```
