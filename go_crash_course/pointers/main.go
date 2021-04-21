package main

import "fmt"

func main() {
	// storing the hexadecimal
	// values in variables
	x := "Hello World"
	y := &x

	// Displaying the values
	fmt.Printf("Value of y in hexdecimal is %X\n", y)
	fmt.Printf("Address of x in hexdecimal is %X\n", &x)
	fmt.Printf("Value that y is point to in hexdecimal is %s\n", *y)

	*y = "dlrow olleh"

	fmt.Printf("x is now: %s", x)

	fmt.Println(&x == y) // prints true because y is the address value of x
}
