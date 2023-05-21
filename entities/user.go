package userdata

const ACTIVE bool = true
const INACTIVE bool = false

type User struct {
	Name   string
	Age    int
	Active bool
}

// constructor for User
// setting default values
// if no values present
func (user *User) FillDefaultUser() {
	if user.Name == "" {
		user.Name = "Default"
	}
	if user.Age == 0 {
		user.Age = 1
	}
	if user.Active == INACTIVE {
		user.Active = ACTIVE
	}
}
