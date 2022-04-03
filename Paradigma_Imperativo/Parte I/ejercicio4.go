package main

import "fmt"

/*
 -----------------  EJERCICIO 4 PARTE 1 -----------------------

- Defina una estructura para un inventario de una tienda que vende zapatos. Cada zapato debe contar con la información
de su modelo (marca), su precio y la talla del mismo que debe ir únicamente del 34 al 44.
*/

type zapato struct {
	marca   string
	modelo  string
	talla   int
	precio  float64
	enStock int
}

func nuevoZapato(marca string, modelo string, talla int, precio float64, stok int) *zapato {
	z := new(zapato)
	if talla >= 34 && talla <= 43 {
		z.marca = marca
		z.modelo = modelo
		z.talla = talla
		z.precio = precio
		z.enStock = stok
		return z
	} else {
		return nil
	}

}

func agregarZapato(zapatos []*zapato, marca string, modelo string, talla int, precio float64, stok int) []*zapato {
	z := nuevoZapato(marca, modelo, talla, precio, stok)
	if z != nil {
		zapatos = append(zapatos, z)
	}
	return zapatos
}

func nuevasTallas(zapatos []*zapato, min int, max int) []*zapato {
	// FUNCION DEL EJERCICIO 4 PARTE 3
	nuevaLista := []*zapato{}
	for _, z := range zapatos {
		if z.talla >= min && z.talla <= max {
			nuevaLista = append(nuevaLista, z)
		}
	}
	return nuevaLista
}

func main() {
	/*
		 -----------------  EJERCICIO 4 PARTE 2 -----------------------

		Escriba un programa que lea la información anterior para 10 zapatos del inventario y los almacene en un arreglo.
	*/

	listaZapatos := []*zapato{}
	listaZapatos = agregarZapato(listaZapatos, "Nike", "Air Max", 32, 15000, 10)
	listaZapatos = agregarZapato(listaZapatos, "Adidas", "T Max", 36, 15000, 60)
	listaZapatos = agregarZapato(listaZapatos, "KY", "Air Max", 38, 15000, 70)
	listaZapatos = agregarZapato(listaZapatos, "BV", "Air Max", 40, 15000, 80)
	listaZapatos = agregarZapato(listaZapatos, "Adidas", "Air Max", 42, 15000, 90)
	listaZapatos = agregarZapato(listaZapatos, "VD", "Air GF", 44, 15000, 100)
	listaZapatos = agregarZapato(listaZapatos, "Adidas", "YTR Max", 42, 15000, 110)
	listaZapatos = agregarZapato(listaZapatos, "YTR", "Air Max", 48, 15000, 120)
	listaZapatos = agregarZapato(listaZapatos, "FDC", "TR Max", 50, 15000, 130)

	fmt.Println("\nEJERCICIO 4 PARTE 2: ")
	for n, z := range listaZapatos {
		fmt.Println((n + 1), z.marca, z.modelo, z.talla, z.precio, z.enStock)
	}

	/*
		 -----------------  EJERCICIO 4 PARTE 3 -----------------------

		Escriba una función que reciba de parámetro dicho arreglo y dos tallas minimo y máximo, y que retorne un nuevo
		arreglo con los zapatos que concuerden con un el rango de tallas entre dichos mínimo y máximo. Documente la
		estrategia utilizada para crear un nuevo conjunto de elementos en tiempo de ejecución para ser retornado por
		la función. OBJETIVO: manejo de estructuras, ciclos, arreglos y slices
		(consideren el uso de slices para este ejercicio)
	*/

	listaZapatos = nuevasTallas(listaZapatos, 36, 38)

	fmt.Println("\nEJERCICIO 4 PARTE 3: ")

	for n, z := range listaZapatos {
		fmt.Println((n + 1), z.marca, z.modelo, z.talla, z.precio, z.enStock)
	}

}
