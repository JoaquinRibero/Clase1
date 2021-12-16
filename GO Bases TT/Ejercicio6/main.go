package main

import ("fmt")

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("La edad de Benjamin es : %d \n", employees["Benjamin"])

	sum := 0
	for _, edad := range employees {
		if edad > 21 {
			sum += 1
		}
	}

	fmt.Printf("Cantidad de empleados mayores a 21 años: %d \n", sum)

	delete(employees, "Pedro")
	fmt.Printf("Se elimina a Pedro del mapa de empleados\n")
	fmt.Println(employees)

}