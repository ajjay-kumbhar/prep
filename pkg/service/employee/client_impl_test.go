package employee

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	coreentity "github.com/ajjay-kumbhar/prep/pkg-core/entity"
	"github.com/ajjay-kumbhar/prep/pkg/service/employee/db"
	"github.com/ajjay-kumbhar/prep/pkg/service/employee/entity"
)

func TestCreateEmployee(t *testing.T) {
	store := db.NewEmployeeStore()
	client := NewClient(store)

	// Create an employee
	emp := &entity.Employee{Name: "John Doe", Position: "Manager", Salary: 50000}
	err := client.CreateEmployee(emp)
	assert.NoError(t, err)
	assert.Equal(t, 1, emp.ID) // First employee should have ID 1

	// Retrieve the employee from the store and check if it matches
	empFromStore, err := client.GetEmployee(1)
	assert.NoError(t, err)
	assert.NotNil(t, empFromStore)
	assert.Equal(t, emp, empFromStore)
}

func TestUpdateEmployee(t *testing.T) {
	store := db.NewEmployeeStore()
	client := NewClient(store)

	// Create an employee
	emp := &entity.Employee{Name: "John Doe", Position: "Manager", Salary: 50000}
	err := client.CreateEmployee(emp)
	assert.NoError(t, err)

	// Update the employee
	updatedEmp := &entity.Employee{ID: 1, Name: "Jane Doe", Position: "Director", Salary: 75000}
	err = client.UpdateEmployee(1, updatedEmp)
	assert.NoError(t, err)

	// Retrieve the employee from the store and check if it matches the updated details
	empFromStore, err := client.GetEmployee(1)
	assert.NoError(t, err)
	assert.NotNil(t, empFromStore)
	assert.Equal(t, updatedEmp, empFromStore)
}

func TestDeleteEmployee(t *testing.T) {
	store := db.NewEmployeeStore()
	client := NewClient(store)

	// Create an employee
	emp := &entity.Employee{Name: "John Doe", Position: "Manager", Salary: 50000}
	err := client.CreateEmployee(emp)
	assert.NoError(t, err)

	// Delete the employee
	err = client.DeleteEmployee(1)
	assert.NoError(t, err)

	// Try to retrieve the deleted employee and expect an error
	_, err = client.GetEmployee(1)
	assert.Error(t, err)
	assert.Equal(t, "employee not found", err.Error())
}

func TestGetEmployees(t *testing.T) {
	store := db.NewEmployeeStore()
	client := NewClient(store)

	// Create some employees
	for i := 0; i < 20; i++ {
		emp := &entity.Employee{Name: fmt.Sprintf("Employee%d", i+1), Position: "Manager", Salary: 50000}
		client.CreateEmployee(emp)
	}

	// Test pagination
	pagination := &coreentity.Pagination{PageNum: 2, PageSize: 10}
	total, employees := client.GetEmployees(pagination)
	assert.Equal(t, 20, total)
	assert.Len(t, employees, 10)
	assert.Equal(t, "Employee11", employees[0].Name) // Check if pagination is correct
}

func TestGetEmployee(t *testing.T) {
	store := db.NewEmployeeStore()
	client := NewClient(store)

	// Create an employee
	emp := &entity.Employee{Name: "John Doe", Position: "Manager", Salary: 50000}
	client.CreateEmployee(emp)

	// Retrieve the employee and check if it matches
	empFromStore, err := client.GetEmployee(1)
	assert.NoError(t, err)
	assert.NotNil(t, empFromStore)
	assert.Equal(t, emp, empFromStore)
}
