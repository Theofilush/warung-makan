package order

import (
	"database/sql"
	"fmt"

	"github.com/Theofilush/warung-makan/model"
)

type OrderRepository interface {
	Create(newOrder model.Order) error
	RetrieveAll() ([]model.Order_details2, error)
	FindById(id string) (model.Order_details2, error)
	Delete(id string) error
	Update(order model.Order) error
}

type orderDbRepository struct {
	db *sql.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "db_warung_makan"
)

func connectDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Connected!")
	}

	return db
}

func validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println("Transaction Rollback")
	} else {
		fmt.Println("Successfully " + message + " data!")
	}
}

func (c *orderDbRepository) Create(newOrder model.Order) error {
	dbb := connectDb()
	tx, err2 := dbb.Begin()
	if err2 != nil {
		panic(err2)
	}

	insertTransaction(newOrder, tx)

	for i := 0; i < len(newOrder.Details.Order_details); i++ {
		insertTransactionDetail(newOrder.Details.Order_details[i], tx)
	}

	err3 := tx.Commit()
	if err3 != nil {
		panic(err3)
	} else {
		fmt.Println("Transaction Commited!")
	}

	return nil
}

func insertTransaction(order model.Order, tx *sql.Tx) {
	insertTransaction := "INSERT INTO m_order VALUES ($1, $2, $3, $4, $5, null, $6)"

	_, err := tx.Exec(insertTransaction, order.Customer_id, order.Table_id, order.Paid_status, order.Total_price, order.Order_detail_id, order.Id)
	validate(err, "Insert order", tx)
}

func insertTransactionDetail(order model.Order_details, tx *sql.Tx) {
	insertTransaction := "INSERT INTO m_order_detail (id, menu_id, quantity, order_id) VALUES ($1, $2, $3, $4)"

	_, err := tx.Exec(insertTransaction, order.Id, order.Menu_id, order.Quantity, order.Order_id)
	validate(err, "Insert Order Detail", tx)
}

func (c *orderDbRepository) RetrieveAll() ([]model.Order_details2, error) {
	rows, err := c.db.Query("select m_order.id, m_order.customer_id, m_order.table_id, m_order.paid_status, m_order.total_price, m_order.order_detail_id, m_order_detail.menu_id, m_order_detail.quantity from m_order join m_order_detail on m_order.order_detail_id = m_order_detail.order_id")
	if err != nil {
		return nil, err
	}
	var orders []model.Order_details2
	for rows.Next() {
		var id string
		var customer_id string
		var table_id int
		var paid_status bool
		var total_price int
		var order_detail_id int
		var menu_id int
		var qty int
		err = rows.Scan(&id, &customer_id, &table_id, &paid_status, &total_price, &order_detail_id, &menu_id, &qty)
		if err != nil {
			return nil, err
		}

		order := model.NewOrder2(id, customer_id, table_id, paid_status, total_price, order_detail_id, menu_id, qty)
		orders = append(orders, order)
	}
	return orders, nil
}

func (c *orderDbRepository) FindById(id string) (model.Order_details2, error) {
	rows, err := c.db.Query("select m_order.id, m_order.customer_id, m_order.table_id, m_order.paid_status, m_order.total_price, m_order.order_detail_id, m_order_detail.menu_id, m_order_detail.quantity from m_order join m_order_detail on m_order.order_detail_id = m_order_detail.order_id where m_order.id=$1", id)
	if err != nil {
		return model.Order_details2{}, err
	}

	var order model.Order_details2
	for rows.Next() {
		var id2 string
		var customer_id string
		var table_id int
		var paid_status bool
		var total_price int
		var order_detail_id int
		var menu_id int
		var qty int
		err = rows.Scan(&id2, &customer_id, &table_id, &paid_status, &total_price, &order_detail_id, &menu_id, &qty)
		if err != nil {
			panic(err)
		}

		order = model.NewOrder2(id2, customer_id, table_id, paid_status, total_price, order_detail_id, menu_id, qty)
	}
	return order, nil
}

func (c *orderDbRepository) Delete(id string) error {
	dbb := connectDb()
	tx, err2 := dbb.Begin()
	if err2 != nil {
		panic(err2)
	}

	ordID := findIdTransaction(id, tx)
	deleteTransaction(id, tx)
	deleteTransactionDetail(ordID, tx)

	err3 := tx.Commit()
	if err3 != nil {
		panic(err3)
	} else {
		fmt.Println("Transaction Commited!")
	}

	return nil
}

func findIdTransaction(id string, tx *sql.Tx) string {
	sumCredit := "SELECT order_detail_id FROM m_order WHERE id = $1"

	var orderDetailID string
	err := tx.QueryRow(sumCredit, id).Scan(&orderDetailID)
	validate(err, "Select Order Detail ID", tx)

	return orderDetailID
}

func deleteTransaction(id string, tx *sql.Tx) {
	deleteTransaction := "delete from m_order where id = $1"

	_, err := tx.Exec(deleteTransaction, id)
	validate(err, "Delete Transsaction", tx)
}

func deleteTransactionDetail(id string, tx *sql.Tx) {
	deleteTransactionDetail := "delete from m_order_detail where order_id = $1"

	_, err := tx.Exec(deleteTransactionDetail, id)
	validate(err, "Delete Transsaction Detail", tx)
}

func (c *orderDbRepository) Update(order model.Order) error {
	updateStatement := "update m_order set customer_name=$2, email=$3, address=$4 where id = $1"
	_, err := c.db.Exec(updateStatement, order.Id)

	if err != nil {
		return err
	}
	return nil
}

// Func seperti consturcor
func NewOrderDbRepository(db *sql.DB) OrderRepository {
	return &orderDbRepository{
		db: db,
	}
}
