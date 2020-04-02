package main

import "fmt"

type User interface {
	permissions() int // 1 - 5
	name() string
}

type Admin struct {
	name string
}

func (this Admin) permissions() int {
	return 5
}

func (this Admin) name() string {
	return ""
}

func auth(user User) string {
	if user.permissions() > 4 {
		return user.name() + " admin"
	}
	return ""
}

func main() {
	admin := Admin{"Uriel"}
	fmt.Println(auth(admin))
}
