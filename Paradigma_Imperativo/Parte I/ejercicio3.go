package main

import (
	"fmt"
	"strings"
)

/*
- Escriba un programa que utilice métodos que reciban como parámetros dos cadenas de caracteres,
y que muestre por pantalla un mensaje que indique si la primera de ellas es una subcadena de la segunda.
Para dicho ejercicio, intente utilizar alguna función predefinida para dicho fin e intente implementar otra versión que
NO utilice ninguna función predefinida. OBJETIVO: uso de librerías, métodos, paso de parámetros, retornos, ciclos.
*/
func main() {
	confunciones("Hola mundo", "mundo")
	sinfunciones("Hola mundo", "gfd")

}

func sinfunciones(cadena string, subCadena string) {
	for i := 0; i <= (len(cadena) - len(subCadena)); i++ {
		if cadena[i:i+len(subCadena)] == subCadena {
			fmt.Println("La subcadena", subCadena, "esta en la cadena", cadena)
			return
		}
	}
	fmt.Println("La subcadena", subCadena, "no esta en la cadena", cadena)
}

func confunciones(cadena string, subCadena string) {
	if strings.Contains(cadena, subCadena) {
		fmt.Println("La cadena", cadena, "contiene la subcadena", subCadena)
	} else {
		fmt.Println("La cadena", cadena, "no contiene la subcadena", subCadena)
	}
}
