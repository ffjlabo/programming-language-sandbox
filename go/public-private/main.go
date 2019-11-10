package main

import (
	"./user"
	"fmt"
)

func main() {
	user := user.NewUser()
	fmt.Println(user)
	fmt.Println(user.Name)
	// fmt.Println(user.age) // privateな値なので怒られる
}

