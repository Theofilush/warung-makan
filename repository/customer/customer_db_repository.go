package customer

import (
	"database/sql"

	"enigmacamp.com/final-project/model"
)

type CustomerRepository interface {
	Create(newCustomer model.Customer) error
	RetrieveAll() ([]model.Customer, error)
	FindById(id string) (model.Customer, error)
	Delete(id string) error
	Update(customer model.Customer) error
}

type customerDbRepository struct {
	db *sql.DB
}

func (c *customerDbRepository) Create(newCustomer model.Customer) error {
	insertStatement := "insert into m_customer values($1,$2,$3,true,$4)"
	_, err := c.db.Exec(insertStatement, newCustomer.Id, newCustomer.Customer_name, newCustomer.Email, newCustomer.Address)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerDbRepository) RetrieveAll() ([]model.Customer, error) {
	rows, err := c.db.Query("select id, customer_name, email, address from m_customer")
	if err != nil {
		return nil, err
	}
	var customers []model.Customer
	for rows.Next() {
		var id string
		var name string
		var email string
		var address string
		err = rows.Scan(&id, &name, &email, &address)
		if err != nil {
			return nil, err
		}

		customer := model.NewCustomer(id, name, email, address)
		customers = append(customers, customer)
	}
	return customers, nil
}

func (c *customerDbRepository) FindById(id string) (model.Customer, error) {
	rows, err := c.db.Query("select id, customer_name, email, address from m_customer where id=$1", id)
	if err != nil {
		return model.Customer{}, err
	}
	var customer model.Customer
	for rows.Next() {
		var id string
		var name string
		var email string
		var address string
		err = rows.Scan(&id, &email, &name, &address)
		if err != nil {
			panic(err)
		}

		customer = model.NewCustomer(id, name, email, address)
	}
	return customer, nil
}

func (c *customerDbRepository) Delete(id string) error {
	deleteStatement := "delete from m_customer where id = $1"
	_, err := c.db.Exec(deleteStatement, id)

	if err != nil {
		return err
	}
	return nil
}

func (c *customerDbRepository) Update(customer model.Customer) error {
	updateStatement := "update m_customer set customer_name=$2, email=$3, address=$4 where id = $1"
	_, err := c.db.Exec(updateStatement, customer.Id, customer.Customer_name, customer.Email, customer.Address)

	if err != nil {
		return err
	}
	return nil
}

// Func seperti consturcor
func NewCustomerDbRepository(db *sql.DB) CustomerRepository {
	return &customerDbRepository{
		db: db,
	}
}
