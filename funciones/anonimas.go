package funciones

import "fmt"

func Operaciones() {
	operacion := func(numero1 int, numero2 int) int {
		return numero1 + numero2
	}

	fmt.Println(operacion(10, 25))
}