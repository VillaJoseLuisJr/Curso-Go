package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*number1, _ := strconv.Atoi(os.Args[1])
	number2, _ := strconv.Atoi(os.Args[2])
	fmt.Println("Sum: ", number1+number2)*/
	sum := sum(os.Args[1], os.Args[2])
	fmt.Println("Sum:", sum)
}

func sum(number1 string, number2 string) (result int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	result = int1 + int2
	return
}
