package main

import (
	"fmt"
)

func main() {
	switch num := 15; {
	case num < 50:
		fmt.Println("%d is less than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Println("%d is greater than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Println("%d is less than 200", num)
	}
}
