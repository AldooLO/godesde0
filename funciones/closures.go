package funciones

import "fmt"

func tabla(numero int) func() int {
	secuencia := 0
	return func () int {
		secuencia++
		return numero * secuencia
	}
}

func LlamarClosure() {
	MiTabla := tabla(2)

	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}
}