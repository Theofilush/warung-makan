package model

import "fmt"

type Order struct {
	Id              string  `json:"id"`
	Customer_id     string  `json:"customer_id"`
	Table_id        int     `json:"table_id"`
	Paid_status     bool    `json:"paid_status"`
	Total_price     int     `json:"total_price"`
	Order_detail_id int     `json:"order_detail_id"`
	Details         Details `json:"details"`
}

type Details struct {
	Order_details []Order_details `json:"order_details"`
}

type Order_details struct {
	Id       int `json:"id"`
	Menu_id  int `json:"menu_id"`
	Quantity int `json:"quantity"`
	Order_id int `json:"order_id"`
}

type Order_details2 struct {
	Id              string `json:"id"`
	Customer_id     string `json:"customer_id"`
	Table_id        int    `json:"table_id"`
	Paid_status     bool   `json:"paid_status"`
	Total_price     int    `json:"total_price"`
	Order_detail_id int    `json:"order_detail_id"`
	Menu_id         int    `json:"menu_id"`
	Quantity        int    `json:"quantity"`
}

func (c Order) String() {
	fmt.Println(c.Id, c.Customer_id, c.Table_id, c.Paid_status, c.Total_price, c.Order_detail_id, c.Details)
}

func (c Order_details2) String2() {
	fmt.Println(c.Id, c.Customer_id, c.Table_id, c.Paid_status, c.Total_price, c.Order_detail_id, c.Menu_id, c.Quantity)
}

func NewOrder(id string, cust_id string, paid_status bool, table_id int, total_price int, order_detail_id int, details Details) Order {
	return Order{
		Id:              id,
		Customer_id:     cust_id,
		Table_id:        table_id,
		Paid_status:     paid_status,
		Total_price:     total_price,
		Order_detail_id: order_detail_id,
		Details:         details,
	}
}

func NewOrder2(id string, cust_id string, table_id int, paid_status bool, total_price int, order_detail_id int, menu_id int, quantity int) Order_details2 {
	return Order_details2{
		Id:              id,
		Customer_id:     cust_id,
		Table_id:        table_id,
		Paid_status:     paid_status,
		Total_price:     total_price,
		Order_detail_id: order_detail_id,
		Menu_id:         menu_id,
		Quantity:        quantity,
	}
}
