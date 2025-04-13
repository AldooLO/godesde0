package main

import (
	"fmt"
	"godesde0/variables"
)

func main() {
	estado, texto := variables.ConvertiraTexto(123)
	fmt.Println(estado)
	fmt.Println(texto)
}