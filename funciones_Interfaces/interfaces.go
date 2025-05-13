package funciones_Interfaces

import (
	"fmt"
	"godesde0/interfaces"
)

func HumanoRespirando(humano interfaces.Humano){
	humano.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", humano.Sexo())
}