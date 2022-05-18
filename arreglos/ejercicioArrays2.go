//Realizar un programa en Golang para imprimir el número más alto de un arreglo de enteros dado.

package arreglos

func MaxNumDeArray(arreglo []int) (numMayor int) {
	listaArreglo := len(arreglo)
	numMayor = 0
	for i := 0; i < listaArreglo; i++ {
		if arreglo[i] > numMayor {
			numMayor = arreglo[i]
		}
	}
	return numMayor
}
