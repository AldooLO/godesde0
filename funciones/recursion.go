package funciones

import "fmt"

func Exponencia(valor int) {
	if valor > 20 {
		return
	}
	fmt.Println(valor)
	Exponencia(valor * 2)
}