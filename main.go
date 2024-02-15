package main

import (
	"fmt"

	"github.com/rocha7778/godesde0/variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConvierteATexto(1588)

	fmt.Println(estado)
	fmt.Println(texto)

}
