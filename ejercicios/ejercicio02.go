package ejercicios

import (
	"fmt"
)

func TablaMultiplicar(number int) {

	fmt.Println("Resultado")

	for i := 1; i <= number; i++ {
		fmt.Print(i)
		fmt.Print(" x ")
		fmt.Print(number)
		fmt.Print(" = ")
		fmt.Print(i * number)
		fmt.Println()

	}

}
