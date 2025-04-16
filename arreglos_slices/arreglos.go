package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{10, 2, 8}

func MostrarArreglos() {
	tabla[2] = 55
	tabla[8] = 23
	fmt.Println(tabla)

	tabla2 := [10]string{"Pablo", "Pedro"}
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}
}