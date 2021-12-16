package main

import "fmt"

func main() {

	var (
		precio float32 = 100
		descuento float32 = 20
	)

	var precioFinal float32 = precio - (precio * (descuento/100))

	fmt.Printf("El precio final con descuento es: %.2f \n", precioFinal)
}