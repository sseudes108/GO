package main

import "fmt"

type Shemales struct {
	name string
	cock int
}

func main() {

	shemales := []Shemales{}
	fmt.Println(shemales)
	alessandra := Shemales{name: "Alessandra Ribeiro", cock: 18}
	carol := Shemales{name: "Carol Penelope", cock: 16}
	fernanda := Shemales{name: "Fernanda Cristine", cock: 19}

	shemales = append(shemales, alessandra, carol, fernanda)
	fmt.Println(shemales)

	for i, shemale := range shemales {
		fmt.Println(i, "Name:", shemale.name, "Tamanho do pau:", shemale.cock)
	}

	/*slice := []string{"Alessandra", "Ribeiro", "Shemale", "Big", "Cock"}

	for i, element := range slice {
		fmt.Println(i, element)
	}*/

	danny := Shemales{name: "Danny Bendochy", cock: 17}
	shemales = append(shemales, danny)

	shemales[3] = shemales[0]
	shemales[0] = danny

	for i, shemale := range shemales {
		fmt.Println(i, "Name:", shemale.name, "Tamanho do pau:", shemale.cock)
	}
}
