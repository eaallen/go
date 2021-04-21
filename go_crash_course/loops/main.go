package main

import (
	"fmt"
	"strconv"
)

func main() {
	// for loop long
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
	// for loop short
	for i := 1; i < 10; i++ {
		fmt.Printf("number %d\n", i)
	}

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		fizz := "Fizz"
		buzz := "Buzz"
		str := ""

		if i%3 == 0 {
			str = str + fizz
		}
		if i%5 == 0 {
			str = str + buzz
		}
		if str == "" {
			str = strconv.Itoa(i)
		}

		fmt.Println(str)

	}
}
