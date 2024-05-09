package db

import "github.com/ajjay-kumbhar/prep/pkg/service/employee/entity"

type EmployeeStore struct {
	Employees []entity.Employee
}

func NewEmployeeStore() *EmployeeStore {
	return &EmployeeStore{
		Employees: []entity.Employee{},
	}
}
