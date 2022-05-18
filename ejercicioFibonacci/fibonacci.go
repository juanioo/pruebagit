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

/*func ImprimiFibonacciSerieArreglo(numTop int) []int {

	result := make([]int, numTop)
	indice := 2

	a := 0
	b := 1
	c := b

	result[0] = a
	result[1] = b

	for true {
		c = b
		b = a + b
		if b >= numTop {
			break
		}
		a = c
		result[indice] = b
		indice++
	}

	return result[:indice]
}

func FibonacciRecursiva(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = FibonacciRecursiva(n-1) + FibonacciRecursiva(n-2)
	}
	return
}*/
