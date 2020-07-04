package main

import (
	"fmt"
	"strconv"
)

type User struct {
	name    string
	age     int64
	gender  bool
	address string
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetAge() int64 {
	return u.age
}

func (u *User) GetGender() bool {
	return u.gender
}

func (u *User) GetAddress() string {
	return u.address
}

func Run7() {
	fmt.Println("\n	7. ")
	user := make(map[string]*User)

	for i := 0; i < 10; i++ {
		obj := User{"name" + strconv.Itoa(i), int64(i), true, "address"}
		user[obj.name] = &obj
		fmt.Println("done ", i)
	}

	for i := range user {
		fmt.Print(user[i].name + "	")
		fmt.Println(user[i].age)
	}

	var copyuser []*User

	for _, value := range user {
		copyuser = append(copyuser, value)
	}

	fmt.Println("Slice copy:")

	for _, value := range copyuser {
		fmt.Print(value.GetName() + " ")
		fmt.Println(value.GetAge())
	}
}
