package main

import (
	"fmt"

	"github.com/DevIvan84/godesde0/ejercicios"
)

func main() {
	//estado, texto := variables.ConviertoaTexti(1566)
	//fmt.Println(estado)
	//fmt.Println(texto)

	//if os := runtime.GOOS; os == "Linux." || os == "OS X." {
	//		fmt.Println("Esto no es Windows")
	//	} else {
	//		fmt.Println("Esto es: ", os)
	//	}

	valor, mensaje := ejercicios.Devuelve2Valores("564")

	fmt.Println(valor)
	fmt.Println(mensaje)

}
