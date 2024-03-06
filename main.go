package main

import (
	"fmt"

	"github.com/DevIvan84/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexti(1566)
	fmt.Println(estado)
	fmt.Println(texto)
}
