package model

import "fmt"

type Suplier_detail struct {
	Id          int `json:"id"`
	Quantity    int `json:"qty"`
	Price_total int `json:"price_total"`
}

func (c Suplier_detail) String() {
	fmt.Println(c.Id, c.Quantity, c.Price_total)
}

func NewSuplier_detail(id, qty, price_total int) Suplier_detail {
	return Suplier_detail{
		Id:          id,
		Quantity:    qty,
		Price_total: price_total,
	}

}
