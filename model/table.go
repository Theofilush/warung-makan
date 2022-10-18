package model

import "fmt"

type Table struct {
	Id                 string `json:"id"`
	Table_number       string `json:"table_number"`
	Reservation_status int    `json:"reservation_status"`
}

func (c Table) String() {
	fmt.Println(c.Id, c.Table_number, c.Reservation_status)
}

func NewTable(id, table_number string, reservation_status int) Table {
	return Table{
		Id:                 id,
		Table_number:       table_number,
		Reservation_status: reservation_status,
	}

}
