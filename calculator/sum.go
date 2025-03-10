package calculator

var logMessage = "[LOG]"

// Version de la calculadora
var Version = "1.0"

func internalSum(number int) int {
	return number - 1
}

// Sum dos numeros enteros
func Sum(number1, number2 int) int {
	return number1 + number2
}
