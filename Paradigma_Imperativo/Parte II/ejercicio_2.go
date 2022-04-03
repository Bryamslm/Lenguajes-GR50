package main

import "fmt"

type persona struct {
	nombre string
	identificador[10] int
}

func masParecido(listaPersonas []persona, muestra[10] int) [3]persona{
	var masParecido[3] persona
	coinsidencias1 := 0
	coinsidencias2 := 0
	coinsidencias3 := 0

	for i := 0; i < len(listaPersonas); i++ {
		p:= listaPersonas[i]
		coinciencias := 0
		for j := 0; j < len(p.identificador); j++ {
			if p.identificador[j] == muestra[j] {
				coinciencias++
			}
		}
		if coinciencias > coinsidencias1 {
			coinsidencias3 = coinsidencias2
			coinsidencias2 = coinsidencias1
			coinsidencias1 = coinciencias
			masParecido[2] = masParecido[1]
			masParecido[1] = masParecido[0]
			masParecido[0] = p
		} else if coinciencias > coinsidencias2 {
			coinsidencias3 = coinsidencias2
			coinsidencias2 = coinciencias
			masParecido[2] = masParecido[1]
			masParecido[1] = p
		} else if coinciencias > coinsidencias3 {
			coinsidencias3 = coinciencias
			masParecido[2] = p
		}
	}

	return masParecido
}

func main() {

	listaPersonas := make([]persona, 0)

	listaPersonas = append(listaPersonas, persona{"Juan", [10]int{1,2,3,4,5,6,7,8,9,10}})
	listaPersonas = append(listaPersonas, persona{"Pedro", [10]int{1,2,3,4,5,6,7,8,10,11}})
	listaPersonas = append(listaPersonas, persona{"Jose", [10]int{1,2,3,4,5,6,10,11,12,13}})
	listaPersonas = append(listaPersonas, persona{"Carlos", [10]int{1,2,3,4,5,10,11,12,13,14}})
	listaPersonas = append(listaPersonas, persona{"Bryam", [10]int{1,2,3,4,10,11,12,13,14,15}})
	listaPersonas = append(listaPersonas, persona{"Bernal", [10]int{1,2,3,10,11,12,13,14,15,16}})
	listaPersonas = append(listaPersonas, persona{"Karla", [10]int{1,2,3,10,11,12,13,14,15,17}})
	listaPersonas = append(listaPersonas, persona{"Maria", [10]int{1,2,3,10,5,6,7,10,11,12}})

	fmt.Println(masParecido(listaPersonas, [10]int{1,2,3,4,5,6,7,8,9,10}))
	fmt.Println(masParecido(listaPersonas, [10]int{1,2,3,10,11,12,13,14,15,20}))
	fmt.Println(masParecido(listaPersonas, [10]int{1,2,3,4,5,6,10,11,12,14}))
}

