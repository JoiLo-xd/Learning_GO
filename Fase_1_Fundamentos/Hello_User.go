package fase1fundamentos

import "fmt"

// HelloWorld es la función que imprime "Hello, World!" en la consola.
func HelloWorld() {
	fmt.Println("Hola, como te llamas?")
	var nombre string
	fmt.Scanln(&nombre)
	fmt.Println("Hola,", nombre, "!")
}
