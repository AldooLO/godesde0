package main

import (
	//"fmt"
	//"godesde0/variables"
	//"godesde0/funciones"
	//"godesde0/modelos"
	//fI "godesde0/funciones_Interfaces"
	"fmt"
	"godesde0/goroutines"
)

func main() {
	/*estado, texto := variables.ConvertiraTexto(123)
	fmt.Println(estado)
	fmt.Println(texto)*/
	/* Pedro := new(modelos.Hombre)
	fI.HumanoRespirando(Pedro)

	Sandra := new(modelos.Mujer)
	fI.HumanoRespirando(Sandra) */
	go goroutines.MiNombreLento("Sergio Perez")

	fmt.Println("Estoy aqui")
}