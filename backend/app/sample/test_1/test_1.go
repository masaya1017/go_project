package test_1

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) User {
	return User{
		Name: name,
		Age:  age,
	}
}

func (u *User) GetMessage(message string) string {
	return fmt.Sprintf("%sさん, %s", u.Name, message)
}
