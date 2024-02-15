package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {

	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()

	fmt.Println("Nombre", Nombre)
	fmt.Println("Estado", Estado)
	fmt.Println("Sueldo", Sueldo)
	fmt.Println("Fecha", Fecha)
	fmt.Println()

}

func ConvierteATexto(numero int) (bool, string) {
	text := strconv.Itoa(numero)
	return true, text
}
