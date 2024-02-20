package funciones

import (
	"fmt"
	"strings"
	"time"
)

func Exponencial(n int) {

	if n > 100_000_000 {
		return
	}

	fmt.Println(n)
	Exponencial(n * 2)

}

func MiNombreLentooo(nombre string, canal1 chan bool) {
	letras := strings.Split(nombre, "")
	for otra_cosa, letra := range letras {

		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("%v otro valor, letra del array %v\n", otra_cosa, letra)

	}
	canal1 <- true
	fmt.Println("Termino")

}
