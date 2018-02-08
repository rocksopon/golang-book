package main

import "fmt"

func main() {
	fmt.Print("Enter a Farenheit: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := ((input - 32) * 5/9)
	fmt.Printf("%.2f Farenheit = %.2f Celsius", input, output)
}