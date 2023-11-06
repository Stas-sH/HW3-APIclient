package main

import (
	"Stas-sH/testclienttomyserver/client"
	"fmt"
	"log"
)

var NewUser client.User = client.User{
	UserName: "Pavlo",
	Mail:     "pavel@mail.com",
	Phone:    "+380111111111",
	Password: "1111",
}

var UpdateUser client.User = client.User{
	UserName: "Tomy",
	Mail:     "tom@mail.com",
	Phone:    "+38022222222",
	Password: "2222",
}

func main() {
	//serverClient, err := client.NewClient(time.Second * 5)
	serverClient, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	//err = serverClient.AddUser(NewUser)
	if err != nil {
		log.Fatal(err)
	}

	//deleteUserId := "3"
	//err = serverClient.DeleteUserById(deleteUserId)
	if err != nil {
		log.Fatal(err)
	}

	updateUserId := "4"
	err = serverClient.UpdateUserById(updateUserId, UpdateUser)
	if err != nil {
		log.Fatal(err)
	}

	users, err := serverClient.GetUsers()
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range users {
		fmt.Println(value.UserInfo())
	}

	idUserForSearch := "2"
	user, err := serverClient.GetUserById(idUserForSearch)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.UserInfo())
}
