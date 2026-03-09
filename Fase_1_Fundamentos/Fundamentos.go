package main

import "fmt"

func main() {
	fmt.Println("Hola, como te llamas?")
	var nombre string
	fmt.Scanln(&nombre) // &nombre es el puntero a la variable nombre, se le tiene que pasar el puntero a la función Scanln.
	fmt.Println("Hola,", nombre, "!")

	// Ejemplos de tipos primitivos
	var edad int       // entero
	var soltero bool   // booleano
	var altura float64 // número con decimales

	var inicial rune   // carácter Unicode (alias de int32)
	var letra byte     // byte (alias de uint8)
	var salario uint   // entero sin signo
	var mensaje string // cadena de texto

	var temperatura float32
	var numeroGrande int64
	var numeroPequeño int8

	// Asignaciones simples para usar las variables
	edad = 30
	soltero = true
	altura = 1.80
	inicial = 'J'
	letra = 'A'
	salario = 1000
	mensaje = "Ejemplo de string"
	temperatura = 36.6
	numeroGrande = 1000000000
	numeroPequeño = -5

	// Imprimir algunos valores
	fmt.Println("Edad:", edad, "Soltero:", soltero)
	fmt.Println("Altura:", altura, "Inicial:", string(inicial), "Letra (byte):", string(letra))
	fmt.Println("Salario:", salario, "Mensaje:", mensaje)
	fmt.Println("Temperatura:", temperatura, "Número grande:", numeroGrande, "Número pequeño:", numeroPequeño)

	// En golang igual que en java podemos darle valores a las variables en la misma linea de codigo.
	var numero int = 10
	fmt.Println("Numero:", numero)

	//Hay una diferencia entre assignacion y declaracion.
	//La declaracion es cuando se crea la variable, la assignacion es cuando se le asigna un valor a la variable.
	//Por ejemplo:
	var numero2 int //declaracion
	numero2 = 10    //assignacion
	fmt.Println("Numero2:", numero2)

	//Tambien podemos declarar las variables y asignarles un valor en la misma linea de codigo, esto se conoce como "short declaration".
	var numero3 int = 10
	fmt.Println("Numero3:", numero3)

	//incluso puede adivinar el tipo de la variable, esto se conoce como "type inference".
	var numero4 = 10
	fmt.Println("Numero4:", numero4)

	//tambien cuando declaramos y asignamos podemos usar el operador :=, esto se conoce como "short assignment".
	numero5 := 10
	fmt.Println("Numero5:", numero5)

	// Tambien podemos declarar las variables y asignarles un valor en la misma linea de codigo, esto se conoce como "short declaration".
	// El compilador de golang es capaz de inferir el tipo de las variables, por lo que no es necesario especificar el tipo de las variables.
	var edad2, soltero2, altura2 = 30, true, 1.80
	fmt.Println("Edad:", edad2, "Soltero:", soltero2, "Altura:", altura2)

	//Conversion de tipos
	//Para convertir un tipo de variable a otro, se puede usar el operador de conversion de tipos.
	//Por ejemplo:
	var numero6 int = 10
	var numero7 float64 = float64(numero6)
	fmt.Println("Numero7:", numero7)

}
