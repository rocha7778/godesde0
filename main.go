package main

import (
	"fmt"

	"github.com/rocha7778/poo/apuntadores"
	"github.com/rocha7778/poo/interfaces"
	"github.com/rocha7778/poo/users"
)

func main() {

	square := &interfaces.Square{Width: 5, Height: 5}
	triangulo := &interfaces.Trinagulo{Base: 4, Altura: 3}

	formas := []interfaces.CalculaArea{square, triangulo}

	// Iterar sobre el array y imprimir el área de cada forma
	for _, forma := range formas {
		area := forma.GetArea()
		nombre := forma.MuestraNombre()
		fmt.Printf("Área del %s: es %.2f\n", nombre, area)
	}

	muestraArea(square)
	muestraArea(triangulo)
	fmt.Println("")

	u := users.AltaUsuarioPorHora()

	fmt.Println(u.IsPaidByhour)
	apuntadores.Apuntadores()

}

func muestraArea(forma interfaces.CalculaArea) {
	fmt.Printf("Área del %s: es %.2f\n", forma.MuestraNombre(), forma.GetArea())
}
