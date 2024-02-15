package main

import (
	"fmt"
	"runtime"

	"github.com/rocha7778/godesde0/variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConvierteATexto(1588)

	fmt.Println(estado)
	fmt.Println(texto)

	os := runtime.GOOS

	if os == "windows" {
		fmt.Println("Windows")
	} else {
		fmt.Println("Linux")
	}

	switch os := runtime.GOOS; os {

	case "windows":
		fmt.Println("Windows")
	case "Linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Other")

	}
}
