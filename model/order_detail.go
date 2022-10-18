package model

import "fmt"

type Order_detail struct {
	Id       int `json:"id"`
	Menu_id  int `json:"menu_id"`
	Quantity int `json:"qty"`
	Order_id int `json:"order_id"`
}

func (c Order_detail) String() {
	fmt.Println(c.Id, c.Menu_id, c.Quantity, c.Order_id)
}

func NewOrder_detail(id, menu_id, qty, order_id int) Order_detail {
	return Order_detail{
		Id:       id,
		Menu_id:  menu_id,
		Quantity: qty,
		Order_id: order_id,
	}

}
