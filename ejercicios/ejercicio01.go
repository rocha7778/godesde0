package ejercicios

import (
	"strconv"
)

func ConversorNumeric(texto string) (int, string) {

	numero, err := strconv.Atoi(texto)

	if err != nil {
		return 0, err.Error()
	}

	if numero > 100 {
		return numero, "El numero es mayor a 100"
	} else {
		return numero, "El numero es menor o igual a 100"
	}

}
