package users

import (
	"github.com/rocha7778/poo/modelos"
)

func AltaUsuario() *modelos.User {
	u := new(modelos.User)

	u.SetAge(10)
	u.SetEmail("rocha7778@gmail.com")
	u.SetId("73216154")
	u.SetName("Rocha")
	u.SetStatus(true)

	return u
}
func AltaUsuarioConstructor(name string, age int, email string, id string, status bool) *modelos.User {
	u := new(modelos.User)
	u.AddUser(name, age, email, id, status)
	return u
}

func AltaUsuarioPorHora() *modelos.UserByHour {

	u := new(modelos.UserByHour)

	u.SetAge(10)
	u.SetEmail("rocha7778@gmail.com")
	u.SetId("73216154")
	u.SetName("Rocha")
	u.SetStatus(true)
	u.SetPaidForHour(true)
	return u

}
