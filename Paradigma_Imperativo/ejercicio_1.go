package main

import "fmt"

func buscar(text string, word string) bool {

	/*
		Función que busca palabra en un texto
		El programa es capaz de diferenciar palabras por los espacios que hay en el texto, no lo hace por
		ejemplo con ; ,
	*/
	var count int
	count = len(text) - len(word)

	for i := 0; i <= count; i++ {
		if text[i:i+len(word)] == word {
			if ((i == 0 && int(text[len(word)]) < 65) || (i == count && int(text[count-1]) < 65) || (int(text[i-1]) < 65 &&
				int(text[i+len(word)]) < 65)) || ((i == 0 && int(text[len(word)]) > 122) || (i == count &&
				int(text[count-1]) > 122) || (int(text[i-1]) > 122 && int(text[i+len(word)]) > 122)) {
				/*
					En el if se comparan los códigos ascii para saber si se encuenta la palabra como correcta, es decir,
					que no sea una palabra a medias o que no corresponde, se aceptan palabras con los signos
					de puntuación habituales acompañandole
				*/
				return true
			}
		}
	}
	return false
}

func main() {

	var text, word string
	text = "A night at the opera. Was good."
	word = "night" //Palabra

	if buscar(text, word) {
		fmt.Println("\n Word was found")
	} else {
		fmt.Println("\n Word was NOT found")
	}
}
