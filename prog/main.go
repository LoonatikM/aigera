package main

import (
	"fmt"
)

func main() {
	var num1 float64
	var num2 float64
	var result float64
	var oper string
	fmt.Print("Enter number 1: ")
	fmt.Scan(&num1)
	fmt.Print("Enter number 2: ")
	fmt.Scan(&num2)
	fmt.Print("What do you want do: ")
	fmt.Scan(&oper)

	if oper == "+" {
		result = num1 + num2
	}
	else oper == "-"(
		result = num1 - num2num1
	)
}
