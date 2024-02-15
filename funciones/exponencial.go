package funciones

import (
	"fmt"
)

func Exponencial(n int) {

	if n > 100_000_000 {
		return
	}

	fmt.Println(n)
	Exponencial(n * 2)

}
