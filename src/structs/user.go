package structs

import "fmt"

// User struct
type User struct {
	ID             int
	Name, Location string
}

// Greetings func
func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}
