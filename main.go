package main

import (
	"fmt"
	"runtime"

	"github.com/rocha7778/godesde0/ejercicios"
	"github.com/rocha7778/godesde0/funciones"
	"github.com/rocha7778/godesde0/mapas"
	"github.com/rocha7778/godesde0/users"
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

	fmt.Println(ejercicios.ConversorNumeric("100"))
	fmt.Println(ejercicios.ConversorNumeric("300"))
	fmt.Println(ejercicios.ConversorNumeric("30a"))
	ejercicios.TablaMultiplicar(10)
	funciones.Calculos()
	funciones.LlamarClosure()
	funciones.Exponencial(1)
	mapas.MostrarMapas()
	fmt.Println(users.AltaUsuario())
	fmt.Println(users.AltaUsuarioConstructor("Rocha", 10, "eldios_delcielo@yahoo.com", "73216154", true))

}
