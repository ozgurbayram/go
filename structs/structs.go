package structs

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	name    string
	surname string
}

func Structs() {
	var ozgur = User{name: "ozgur", surname: "Bayram"}

	ozgur.capitalizeValues()

	fmt.Println(ozgur.name)
	fmt.Println(ozgur.surname)

	fmt.Println("Uset struct is",ozgur.toString())
}

func (user *User) capitalizeValues() {
	var users = reflect.ValueOf(user)

	fmt.Println("dsa", users)
	user.name = strings.ToUpper(user.name)
}

func (user User) toString() string {
	return fmt.Sprintf("%s %s", user.name, user.surname)
}
