package mapas

import (
	"fmt"
)

func MostrarMapas() {
	mapa := make(map[string]string)
	mapa["nombre"] = "Luis"
	mapa["edad"] = "20"
	mapa["pais"] = "Colombia"
	fmt.Println(mapa)
	fmt.Println(mapa["edad"])
	fmt.Println(mapa["pais"])

	campeonato := map[string]int{
		"1":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
	}

	fmt.Println(campeonato)
}
