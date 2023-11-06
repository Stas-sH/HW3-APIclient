package client

import "fmt"

type UsersResponse struct {
	Users []User
}

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"name"`
	Mail     string `json:"mail"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (u User) UserInfo() string {
	return fmt.Sprintf("[%d]-id | Name - %s | Mail - %s | Phone - %s | Password - %s ;\n", u.Id, u.UserName, u.Mail, u.Phone, u.Password)
}
