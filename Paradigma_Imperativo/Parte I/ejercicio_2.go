package main

import "fmt"

/*
- Escriba un programa que inserte un elemento en una determinada posición de un arreglo mediante el uso de punteros.
	El ejercicio debe contemplar la perdida mediante sustitución del elemento existente en dicho arreglo para la
	posición seleccionada. Una vez terminado el ejercicio, intente implementar otra variante del mismo que permita
	no perder el elemento sustituido, si no ampliar el tamaño original del arreglo. Mencione sus allazgos mediante
	documentación interna. OBJETIVO: Uso de arreglos y punteros e investigación para tratar de ampliar de forma dinámica
	el tamaño de un arreglo
*/

func sustituir(element *string, posicion int, array *[5]string) {

	for i := range array {
		if posicion == i {
			array[posicion] = *element
			/*
				En mi mente esto funciona así: solo hago que el puntero del array en la posición i apunte hacia donde está
				apuntando el puntero *element, esto funciona ya que esta función no retorna nada pero aún así el array cuando se
				imprime se ve alterado con el nuevo elemento insertado.
			*/
			break
		}
	}
}

func sustitucion() {
	array := [5]string{"Bryam", "Alejandro", "Ariana", "Alberto", "Maria"}
	element := "Karla"             //elemento a agregar
	sustituir(&element, 1, &array) //posición es donde se va a sustituir, se trabaja con posición de indice array
	fmt.Println(array)
}

func agregar(element *int, posicion int, array2 []int) {

	/*
		Hallasgos: De momento no he sido capaz de acceder a los valores del slice mediante el paso del slice por referencia
		ya que al no especificar el tamaño del slice el compilador no es capaz de saber los índices existentes. Si especifico
		el tamaño del slice en el parámetro no zoy capaz de enviar el slice por referencia. De momento hice redundancia al
		poner *& al inicio del slice y variables. Se tratará de corregir medianre reunión con el profesor o tutor.
	*/

	temp1 := 0
	temp2 := 0
	for i, i2 := range *&array2 {
		if i < posicion {
			continue
		} else if i == posicion {
			temp1 = i2
			*&array2[i] = *element
			if i == len(array2)-1 {
				*&array2 = append(*&array2, temp1)
			}
		} else {
			temp2 = i2
			*&array2[i] = *&temp1
			temp1 = temp2
			if i == len(*&array2)-1 {
				*&array2 = append(*&array2, *&temp1)
			}
		}

	}

	fmt.Println(array2)
}
func sinSustitucion() {
	array := make([]int, 4, 10)
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 5
	element := 4 //elemento a agregar
	fmt.Println(array)
	agregar(&element, 3, array) //posición es donde se va a sustituir, se trabaja con posición de indice array
	//fmt.Println(array)
}

func main() {
	sustitucion()
	sinSustitucion()
}
