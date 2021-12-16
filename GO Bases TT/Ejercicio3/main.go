package main

import ("fmt")

func main() {

	var (
		edad int = 25
		empleado bool = true
		antiguedad int = 2
		sueldo int = 90000
	)

	if (edad > 22 && empleado == true && antiguedad > 1) {
		if (sueldo > 100000) {
			fmt.Println("Se le otorga el prestamo sin intereses.")
		} else {
			fmt.Println("Se le otorga el prestamo con intereses.")
		}
	} else {
		fmt.Println("No se le otorga prestamo.")
	}

}