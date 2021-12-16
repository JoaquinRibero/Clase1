package main

import "fmt"

func main() {

	var palabra string = "academia"

	n := len(palabra);

	fmt.Printf("La cantidad de letras de la palabra \"%s\" es: %d\n", palabra, n)

	for _, value := range palabra {
		fmt.Printf("%c \n", value)
	}
}