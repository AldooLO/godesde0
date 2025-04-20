package arreglos_slices

import "fmt"

var slices []int = []int{1, 2, 3}
var arreglo [10]int = [10]int{12,3,76,44,5,1,34,6}
func MostrarSlices() {
	fmt.Println(slices)
	fmt.Println(arreglo[3:])
	fmt.Println(arreglo[:5])
	fmt.Println(arreglo[6:7])
}

func Capacidad(){
	elementos := make([]int, 5, 20)  //creacion de un slice que crea el tamaño y capacidad se reserva en memoria para performan
	//duplica la capacidad cada vez que se queda sin espacio
	fmt.Printf("tamaño %d, capacidad %d", len(elementos), cap(elementos))

	numeros := make([]int, 0, 0)
	for i := 0; i < 100; i++{
		numeros = append(numeros, i)
	}
	fmt.Printf("\ntamaño %d, capacidad %d", len(numeros), cap(numeros)) // potencia de 2
}