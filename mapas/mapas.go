package mapas

import "fmt"

func MostrarMapas() {
	/* paises := make(map[string]string)
	paises["México"] = "CDMX"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["México"]) */

	edades := map[string]int{
		"Felipe": 39,
		"Mario": 20,
		"Pedro": 10,
	}

	/* for nombre, edad := range edades{
		fmt.Printf("Nombre %s, su edad es %d \n", nombre, edad)
	} */

	delete(edades, "Mario")
	fmt.Println(edades)

	edad, existe := edades["Manolo"]
	fmt.Printf("Su edad es %d, y el nombre existe = %t \n",edad, existe)
}

