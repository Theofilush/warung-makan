package model

import "fmt"

type Reservation struct {
	Id               string `json:"id"`
	Reservation_date string `json:"reservation_date"`
	Table_number     int    `json:"table_number"`
	Customer_id      string `json:"customer_id"`
}

func (c Reservation) String() {
	fmt.Println(c.Id, c.Reservation_date, c.Table_number, c.Customer_id)
}

func NewReservation(id, reservation_date, customer_id string, table_number int) Reservation {
	return Reservation{
		Id:               id,
		Reservation_date: reservation_date,
		Table_number:     table_number,
		Customer_id:      customer_id,
	}

}
