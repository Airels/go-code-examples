package main

import (
	"fmt"
)

type User struct {
	Username string
	IsAdmin bool
}


func main() {
	var airels User
	airels.Username = "Airels"
	airels.IsAdmin = true

	var root User
	root = User{"Root", true}

	moka := User{"Mokapus", false}

	var lamboulda User
	lamboulda = User {
		Username: "Lambda_Aur",
		IsAdmin: false,
	}

	var users [4]User
	users[0] = airels
	users[1] = root
	users[2] = moka
	users[3] = lamboulda

	for _, element := range users { // "_" means ignoring index (not used)
		fmt.Println("User:", element.Username, "\t Admin: ", element.IsAdmin)
	}
}
