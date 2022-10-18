package model

import "fmt"

type Employee struct {
	Id            string `json:"id"`
	Employee_name string `json:"name"`
	Email         string `json:"email"`
	Address       string `json:"address"`
	Phone_number  string `json:"phone_number"`
	Role_employee string `json:"role_employee"`
}

func (c Employee) String() {
	fmt.Println(c.Id, c.Employee_name, c.Email, c.Address, c.Phone_number, c.Role_employee)
}

func NewEmployee(id, name, email, address, phone_number, role_employee string) Employee {
	return Employee{
		Id:            id,
		Employee_name: name,
		Email:         email,
		Address:       address,
		Phone_number:  phone_number,
		Role_employee: role_employee,
	}

}
