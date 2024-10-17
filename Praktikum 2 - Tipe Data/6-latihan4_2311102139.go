package main

import (
	"fmt"
)

func latihan4() {
	var celcius, fahrenheit, kelvin, reamur float64

	fmt.Print("Temperatur Celsius: ")
	fmt.Scanln(&celcius)

	reamur = celcius * 4 / 5
	fahrenheit = (celcius * 9 / 5) + 32
	kelvin = celcius + 273.15
	
	fmt.Println("Derajat Reamur: ", reamur)
	fmt.Println("Derajat Fahrenheit: ", fahrenheit)
	fmt.Println("Derajat Kelvin: ", kelvin)
}