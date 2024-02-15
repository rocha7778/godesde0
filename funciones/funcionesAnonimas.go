package funciones

import "fmt"

func Calculos() {

	calculo := func(numer1 int, numero2 int) int {
		return numer1 + numero2
	}
	fmt.Println(calculo(2, 3))

	calculo = func(numero1 int, numero2 int) int {
		return numero1 * numero2
	}
	fmt.Println(calculo(2, 3))
}
