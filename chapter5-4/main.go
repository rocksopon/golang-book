package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var input int
	rand.Seed(time.Now().UTC().UnixNano())
	random := rand.Intn(100)
	fmt.Println(random)
	
	for i := 0 ; i <= 5 ; i++ {
		fmt.Print("Enter Number: ")
		fmt.Scanf("%v\n", &input)
		if i<5 {
			switch {
			case input == random:
				fmt.Println("เจอเเล้ว")
			case input > random:
				fmt.Println("มากไป")
			case input < random:
				fmt.Println("น้อยไป")
			}
		}else {
			fmt.Println("เกินพอ")
		}
		}
		
}
