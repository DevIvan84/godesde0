package ejercicios

import (
	"strconv"
)

func Devuelve2Valores(texto string) (int, string) {

	entero, err := strconv.Atoi(texto)

	if err != nil {
		return 0, "hubo un error" + err.Error()
	}
	if entero > 100 {
		return entero, "es mayor a 100"
	} else {
		return entero, "es menor a 100"
	}

}
