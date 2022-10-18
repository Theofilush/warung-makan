package model

import "fmt"

type Customer struct {
	Id            string `json:"id"`
	Customer_name string `json:"name"`
	Email         string `json:"email"`
	Address       string `json:"address"`
}

func (c Customer) String() {
	fmt.Println(c.Id, c.Customer_name, c.Email, c.Address)
}

func NewCustomer(id, name, email, address string) Customer {
	return Customer{
		Id:            id,
		Customer_name: name,
		Email:         email,
		Address:       address,
	}

}
