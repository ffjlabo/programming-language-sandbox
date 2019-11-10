package user

type User struct {
	Name string
	age int
}

func NewUser() *User {
	return &User{
		Name: "Example", // public
		age: 22,         // private
	}
}

// private
func getAge(user User) int {
	return user.age
}

