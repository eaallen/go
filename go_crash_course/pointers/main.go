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

	// change the value that y is pointing to
	*y = "dlrow olleh"

	fmt.Printf("x is now: %s \n", x) // x = "dlrow olleh"

	fmt.Println(&x == y) // prints true because y is the address value of x
}
