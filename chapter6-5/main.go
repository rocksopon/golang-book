package main

import (

	"fmt"
)

func main() {
	
	for number := 1 ; number <=100 ; number++ {
		fmt.Println(fizzBuzz(number))
	}
}

func fizzBuzz(number int) (int,string) {
	if number%15 == 0 {
		return  number," is FizzBuzz"
	}else if number%3 == 0 {
		return  number," is Fizz"
	}else if number%5 == 0 {
		return  number," is Buzz"
	}else {
		return number," is not Fizz, Buzz or FizzBuzz"
	}
}
