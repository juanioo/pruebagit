//Crear un programa en Golang
//para imprimir los 100 primeros números de la sucesión de Fibonacci.

package ejercicioFibonacci

import "fmt"

func Fibonacci(cantImprimir int) {
	var a, b, aux int
	a = 0
	b = 1
	aux = b
	fmt.Printf("Serie: %d %d", a, b)
	for true {
		aux = b
		b = a + b
		if b >= cantImprimir {
			fmt.Println()
			break
		}
		a = aux
		fmt.Printf(" %d", b)
	}
}
