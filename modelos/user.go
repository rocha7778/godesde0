package modelos

type User struct {
	Name   string
	Age    int
	Email  string
	Id     string
	Status bool
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) GetAge() int {
	return u.Age
}

func (u *User) SetAge(age int) {
	u.Age = age
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) GetId() string {
	return u.Id
}

func (u *User) SetId(id string) {
	u.Id = id
}

func (u *User) GetStatus() bool {
	return u.Status
}

func (u *User) SetStatus(status bool) {
	u.Status = status
}

func (u *User) AddUser(Name string, Age int, Email string, Id string, Status bool) {
	u.Name = Name
	u.Age = Age
	u.Email = Email
	u.Id = Id
	u.Status = Status
}
