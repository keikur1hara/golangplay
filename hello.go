package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
	//	X,Y int
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.nNm
}

func main() {
	var user1 User
	user1.Age = 29
	user1.Name = "Kei"
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	user4 := User{"user4", 40}
	fmt.Println(user4)

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	user7 := new(User)
	fmt.Println(user7)

	user8 := &User{}
	fmt.Println(user8)

}
