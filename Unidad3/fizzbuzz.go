package main

import "fmt"

func main() {
	for num := 1; num <= 100; num++ {
		switch {
		case num%3 == 0 && num%5 == 0:
			fmt.Println("FizzBuzz")
		case num%3 == 0:
			fmt.Println("Fizz")
		case num%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println("El número es: ", num)
		}
	}
}
