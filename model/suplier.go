package model

import "fmt"

type Suplier struct {
	Id                string `json:"id"`
	Suplier_name      string `json:"name"`
	Faktur_number     string `json:"email"`
	Receive_date      string `json:"receive_date"`
	Id_food_stuff     string `json:"id_food_stuff"`
	Price_total       int    `json:"price_total"`
	Employee_receiver int    `json:"employee_receiver"`
}

func (c Suplier) String() {
	fmt.Println(c.Id, c.Suplier_name, c.Faktur_number, c.Receive_date, c.Id_food_stuff, c.Price_total, c.Employee_receiver)
}

func NewSuplier(id, suplier_name, faktur_number, receive_date, id_food_stuff string, price_total, employee_receiver int) Suplier {
	return Suplier{
		Id:                id,
		Suplier_name:      suplier_name,
		Faktur_number:     faktur_number,
		Receive_date:      receive_date,
		Id_food_stuff:     id_food_stuff,
		Price_total:       price_total,
		Employee_receiver: employee_receiver,
	}
}
