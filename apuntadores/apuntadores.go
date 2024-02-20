package apuntadores

import "fmt"

func Apuntadores() {
	// Declaración de una variable y asignación de un valor
	x := 10

	// Declaración de un apuntador que apunta a la dirección de memoria de x
	//var p *int
	p := &x

	// Modificación del valor de x a través del apuntador p
	*p = 20

	fmt.Println("Valor de x:", x) // Output: 20
}
