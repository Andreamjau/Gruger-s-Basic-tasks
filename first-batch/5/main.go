package main

import (
	"fmt"
)

//write a program that converts celsius to fahrenheit (ggl for equation)


func main(){
	var input int
	var temp float32

	fmt.Println("Enter number for selection\n1. Celsius to fahrenheit\n2. Fahrenheit to celsius");
	fmt.Scan(&input)

	fmt.Println("Enter temperature that should be converted, numbers only")
	fmt.Scan(&temp)

	switch input {
	case 1:
		celsiusTofFahrenheit(temp)
	case 2:
		fahrenheitToCelsius(temp)
	default:
		fmt.Println("How can you mess up typig 1 or 2")
		return
	}
}

func celsiusTofFahrenheit(temp float32){
	result := (temp*1.8) + 32
	fmt.Print(temp, "c is ", result, " degres in Fahrenheit")
}

func fahrenheitToCelsius(temp float32){
	result := (temp - 32)/1.8
	fmt.Print(temp, "f is ", result, " degres in Celsius")
}