package model

import "fmt"

type Order struct {
	Id              string `json:"id"`
	Customer_id     string `json:"name"`
	Table_id        int    `json:"table_id"`
	Paid_status     string `json:"paid_status"`
	Total_price     int    `json:"total_price"`
	Order_detail_id int    `json:"order_detail_id"`
}

func (c Order) String() {
	fmt.Println(c.Id, c.Customer_id, c.Table_id, c.Paid_status, c.Total_price, c.Order_detail_id)
}

func NewOrder(id, cust_id, paid_status string, table_id, total_price, order_detail_id int) Order {
	return Order{
		Id:              id,
		Customer_id:     cust_id,
		Table_id:        table_id,
		Paid_status:     paid_status,
		Total_price:     total_price,
		Order_detail_id: order_detail_id,
	}

}
