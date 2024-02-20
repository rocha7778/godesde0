package modelos

type UserByHour struct {
	IsPaidByhour bool
	User
}

func (u *UserByHour) SetPaidForHour(status bool) {
	u.IsPaidByhour = status
}
