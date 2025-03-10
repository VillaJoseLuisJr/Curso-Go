package main

import (
	"fmt"
	"strconv"
)

/*var defaultInt int
var defaultFloat32 float32
var defaultFloat64 float64
var defaultBool bool
var defaultString string

var integer16 int16 = 127
var integer32 int32 = 32767*/

func main() {
	/*fullName := "John Doe \t(alias \"Foo\")\n"
	fmt.Println(fullName)
	fmt.Println(defaultInt, defaultFloat32, defaultFloat64, defaultBool, defaultString)
	fmt.Println(int32(integer16) + integer32)*/
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, s)
}
