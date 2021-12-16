package main

import "fmt"

func main() {
	mes := 3

	switch mes {
		case 1:
			fmt.Printf("%d, Enero \n", mes)
		case 2:
			fmt.Printf("%d, Febrero \n", mes)
		case 3:
			fmt.Printf("%d, Marzo \n", mes)
		case 4:
			fmt.Printf("%d, Abril \n", mes)
		case 5:
			fmt.Printf("%d, Mayo \n", mes)
		case 6:
			fmt.Printf("%d, Junio \n", mes)
		case 7:
			fmt.Printf("%d, Julio \n", mes)
		case 8:
			fmt.Printf("%d, Agosto \n", mes)
		case 9:
			fmt.Printf("%d, Septiembre \n", mes)
		case 10:
			fmt.Printf("%d, Octubre \n", mes)
		case 11:
			fmt.Printf("%d, Noviembre \n", mes)
		case 12:
			fmt.Printf("%d, Diciembre \n", mes)
		default: 
			fmt.Printf("El número proporcionado no corresponde a un mes")

		//Este ejercicio se podrìa resolver con:
		// 1: Un If por cada valor
		// 2: Un Map en el que la clave es el numero y el valor el mes o vicesversa
		// 3: Utilizando una modulo de fechas.
	}
}