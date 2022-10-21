package model

import (
	"database/sql"
	"fmt"
)

type Menu struct {
	Id        string         `json:"id"`
	Menu_name string         `json:"menu_name"`
	Price     int            `json:"price"`
	Image     sql.NullString `json:"images"`
}

type Menu2 struct {
	Image sql.NullString `json:"images"`
}

func (c Menu) String() {
	fmt.Println(c.Id, c.Menu_name, c.Price, c.Image)
}

// Body:sql.NullString{String:"", Valid:false},
// User:sql.NullInt64{Int64:0, Valid:false}}
func NewMenu(id, name string, image sql.NullString, price int) Menu {
	return Menu{
		Id:        id,
		Menu_name: name,
		Price:     price,
		Image:     image,
	}

}

func NewMenu2(image sql.NullString) Menu2 {
	return Menu2{
		Image: image,
	}

}
