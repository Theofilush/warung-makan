package customer

import (
	"database/sql"
	"errors"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Theofilush/warung-makan/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dummyCustomers = []model.Customer{
	{
		Id:            "1",
		Customer_name: "Dummy Name 1",
		Email:         "Email 1",
		Address:       "Dummy Address 1",
	},
	{
		Id:            "2",
		Customer_name: "Dummy Name 2",
		Email:         "Email 1",
		Address:       "Dummy Address 2",
	},
}

type CustomerRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
}

func (suite *CustomerRepositoryTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	if err != nil {
		log.Fatalln("An error when opening a database connection")
	}
	suite.mockDb = mockDb
	suite.mockSql = mockSql
}

func (suite *CustomerRepositoryTestSuite) TearDownTest() {
	suite.mockDb.Close()
}

func (suite *CustomerRepositoryTestSuite) TestCustomerFindById_Success() {
	dummyCustomer := dummyCustomers[0]
	rows := sqlmock.NewRows([]string{"id", "nama", "address"})
	rows.AddRow(dummyCustomer.Id, dummyCustomer.Customer_name, dummyCustomer.Email, dummyCustomer.Address)
	// buat query mock nya (menggunakan regex -> (.+)
	suite.mockSql.ExpectQuery("select \\* from m_customer where id").WillReturnRows(rows)

	// panggil repository aslinya
	repo := NewCustomerDbRepository(suite.mockDb)

	// panggil method yang mau dtest
	actual, err := repo.FindById(dummyCustomer.Id)

	// buat test assertion
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), actual)
}

func (suite *CustomerRepositoryTestSuite) TestCustomerFindById_Failed() {
	dummyCustomer := dummyCustomers[0]
	rows := sqlmock.NewRows([]string{"ids", "namaaaa", "addresssss"})
	rows.AddRow(dummyCustomer.Id, dummyCustomer.Customer_name, dummyCustomer.Email, dummyCustomer.Address)
	// buat query mock nya (menggunakan regex -> (.+)
	suite.mockSql.ExpectQuery("select \\* from m_customer where id").WillReturnError(errors.New("failed"))

	// panggil repository aslinya
	repo := NewCustomerDbRepository(suite.mockDb)

	// panggil method yang mau dtest
	actual, err := repo.FindById(dummyCustomer.Id)

	// buat test assertion
	func() {
		defer func() {
			if r := recover(); r == nil {
				assert.Error(suite.T(), err)
			}
		}()
		// This function should cause a panic
		repo.FindById(dummyCustomer.Id)
	}()
	assert.NotEqual(suite.T(), dummyCustomer, actual)
	assert.Error(suite.T(), err)
}

func (suite *CustomerRepositoryTestSuite) TestCustomerCreate_Success() {
	dummyCustomer := dummyCustomers[0]
	suite.mockSql.ExpectExec("insert into m_customer values").WithArgs(dummyCustomer.Id, dummyCustomer.Customer_name, dummyCustomer.Email, dummyCustomer.Address).WillReturnResult(sqlmock.NewResult(1, 1))
	repo := NewCustomerDbRepository(suite.mockDb)
	resultRepo := repo.Create(dummyCustomer)
	assert.Nil(suite.T(), resultRepo)
}

func (suite *CustomerRepositoryTestSuite) TestCustomerCreate_Failed() {
	dummyCustomer := dummyCustomers[0]
	suite.mockSql.ExpectExec("insert into m_customer values").WillReturnError(errors.New("failed"))
	repo := NewCustomerDbRepository(suite.mockDb)
	err := repo.Create(dummyCustomer)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), errors.New("failed"), err)
}

func (suite *CustomerRepositoryTestSuite) TestCustomerRetrieveAll_Success() {
	rows := sqlmock.NewRows([]string{"id", "customer_name", "email", "address"})
	for _, v := range dummyCustomers {
		rows.AddRow(v.Id, v.Customer_name, v.Email, v.Address)
	}
	suite.mockSql.ExpectQuery("select \\* from m_customer").WillReturnRows(rows)
	repo := NewCustomerDbRepository(suite.mockDb)

	actual, err := repo.RetrieveAll()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(actual))
	assert.Equal(suite.T(), "C001", actual[0].Id)

}

func (suite *CustomerRepositoryTestSuite) TestCustomerRetrieveAll_Failed() {
	// siapkan column (sama seperti di field table customer)
	rows := sqlmock.NewRows([]string{"ids", "namaaaaa", "addresssss"})
	for _, v := range dummyCustomers {
		rows.AddRow(v.Id, v.Customer_name, v.Email, v.Address)
	}

	// buat query mock nya (menggunakan regex -> (.+)
	suite.mockSql.ExpectQuery("select \\* from m_customer").WillReturnError(errors.New("failed"))

	// panggil repository aslinya
	repo := NewCustomerDbRepository(suite.mockDb)

	// panggil method yang mau dtest
	actual, err := repo.RetrieveAll()

	// buat test assertion
	assert.Nil(suite.T(), actual)
	assert.Error(suite.T(), err)
}

func (suite *CustomerRepositoryTestSuite) TestCustomerDelete_Success() {
	dummyCustomer := dummyCustomers[0]
	suite.mockSql.ExpectExec("delete from m_customer where id = $1").WithArgs("1").WillReturnResult(sqlmock.NewResult(1, 1))
	repo := NewCustomerDbRepository(suite.mockDb)
	resultRepo := repo.Delete(dummyCustomer.Id)
	assert.Nil(suite.T(), resultRepo)
}

func (suite *CustomerRepositoryTestSuite) TestCustomerDelete_Failed() {
	dummyCustomer := dummyCustomers[0]
	suite.mockSql.ExpectExec("delete from m_customer where id = $1").WillReturnError(errors.New("failed"))
	repo := NewCustomerDbRepository(suite.mockDb)
	err := repo.Delete(dummyCustomer.Id)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), errors.New("failed"), err)
}

func TestCustomerRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepositoryTestSuite))
}
