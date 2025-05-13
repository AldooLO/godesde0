package main

import (
	//"fmt"
	//"godesde0/variables"
	//"godesde0/funciones"
	//"godesde0/modelos"
	//fI "godesde0/funciones_Interfaces"
	"godesde0/webserver"
)

func main() {
	/*estado, texto := variables.ConvertiraTexto(123)
	fmt.Println(estado)
	fmt.Println(texto)*/
	/* Pedro := new(modelos.Hombre)
	fI.HumanoRespirando(Pedro)

	Sandra := new(modelos.Mujer)
	fI.HumanoRespirando(Sandra) */

	/* canal1 := make(chan bool)
	go goroutines.MiNombreLento("Sergio Perez", canal1)
	defer func(){
		<-canal1
	}()
	fmt.Println("Estoy aqui") */
	webserver.MiWebServer()
}