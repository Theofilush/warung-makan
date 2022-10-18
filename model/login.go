package model

import "fmt"

type Login struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Id_customer string `json:"id_customer"`
}

func (c Login) String() {
	fmt.Println(c.Id, c.Username, c.Password, c.Id_customer)
}

func NewLogin(id, username, password, id_customer string) Login {
	return Login{
		Id:          id,
		Username:    username,
		Password:    password,
		Id_customer: id_customer,
	}

}
