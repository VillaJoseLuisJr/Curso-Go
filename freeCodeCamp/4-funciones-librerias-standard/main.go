package main

import "fmt"

var funciones = map[string]func(int, int) int{
	"suma":           func(a int, b int) int { return a + b },
	"resta":          func(a int, b int) int { return a - b },
	"multiplicación": func(a int, b int) int { return a * b },
	"división":       func(a int, b int) int { return a / b },
}

func presentaResultado(operacion string, a int, b int) {
	f, exists := funciones[operacion]

	if !exists {
		fmt.Println("Operación no valida")
		return
	}

	fmt.Println("Para a = ", a, " y b = ", b, " la ", operacion, " resultante es: ", f(a, b))
}

func main() {
	presentaResultado("suma", 27, 96)
	presentaResultado("resta", 150, 69)
	presentaResultado("multiplicación", 16, 4)
	presentaResultado("división", 150, 5)
	presentaResultado("multiplica", 16, 4)

}
