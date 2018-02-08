package main

import "fmt"

func main() {
	var x string = "Hello World"
	fmt.Println(x)
	//fmt.Printf("Type: %T\n", x)

	var y string
	y = "Hello World"
	fmt.Println(y)
	//fmt.Printf("Type: %T\n", y)

	z := "Hello World"
	fmt.Println(z)
	fmt.Printf("Type: %T\n", z)

	zz := 123
	fmt.Println(zz)
	fmt.Printf("Type: %T\n", zz)

	const zzz string ="Hello World"

	var (
		a = 5
		b = 6
		c = 7
	)

	fmt.Println(a, b, c)

	v1, v2 := "1st", "2nd"
	fmt.Println(v1, v2)

	v1, v2 = v2, v1
	fmt.Println(v1, v2)



}

