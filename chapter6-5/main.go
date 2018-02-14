package main

import (
	"strconv"
	"fmt"
)

func main() {
	
	for number := 1 ; number <=100 ; number++ {
		fmt.Println(number, fizzBuzz(number))
	}
}

func fizzBuzz(number int) string {
	var str string
	if number%15 == 0 {
		str =  "is FizzBuzz"
	}else if number%3 == 0 {
		str =  "is Fizz"
	}else if number%5 == 0 {
		str =  "is Buzz"
	}else {
		str =  strconv.Itoa(number)
	}
	return str
}
