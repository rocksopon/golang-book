package main

import (
	"strconv"
	"fmt"
)

func main() {
	weatherCelcius(25, "Bangkok few cloud")
}

func weatherCelcius(celcius int, description string) {
	two :=  ` _
 _|
|_`
	
	five := ` _
|_
 _|`
 
	celciusStr := strconv.Itoa(celcius)
	for _, value := range celciusStr {
		valueStr := string(value)
		if valueStr == "2" {
			fmt.Println(two)
		}else if  valueStr == "5" {
			fmt.Println(five)
		}
	}
	fmt.Println(description)
}

