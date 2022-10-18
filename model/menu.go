package model

import "fmt"

type Menu_name struct {
	Id        string `json:"id"`
	Menu_name string `json:"menu_name"`
	Price     int    `json:"price"`
}

func (c Menu_name) String() {
	fmt.Println(c.Id, c.Menu_name, c.Price)
}

func NewMenu_name(id, name string, price int) Menu_name {
	return Menu_name{
		Id:        id,
		Menu_name: name,
		Price:     price,
	}

}
