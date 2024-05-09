package employee

import (
	"fmt"
	"sync"

	coreentity "github.com/ajjay-kumbhar/prep/pkg-core/entity"
	"github.com/ajjay-kumbhar/prep/pkg/service/employee/db"
	"github.com/ajjay-kumbhar/prep/pkg/service/employee/entity"
)

type clientImpl struct {
	sync.RWMutex
	store *db.EmployeeStore
	index int
}

func (c *clientImpl) CreateEmployee(employee *entity.Employee) error {
	c.Lock()
	defer c.Unlock()

	employee.ID = c.index
	c.index = c.index + 1
	c.store.Employees = append(c.store.Employees, *employee)
	return nil
}

func (c *clientImpl) findEmployeeIndex(empID int) int {
	for i := range c.store.Employees {
		if c.store.Employees[i].ID == empID {
			return i
		}
	}
	return -1
}

func (c *clientImpl) UpdateEmployee(empID int, employee *entity.Employee) error {
	c.Lock()
	defer c.Unlock()

	found := c.findEmployeeIndex(empID)
	if found == -1 {
		return fmt.Errorf("employee not found")
	}

	c.store.Employees[found] = *employee

	return nil
}

func (c *clientImpl) DeleteEmployee(empID int) error {
	c.Lock()
	defer c.Unlock()

	found := c.findEmployeeIndex(empID)
	if found == -1 {
		return fmt.Errorf("employee not found")
	}

	c.store.Employees = append(c.store.Employees[:found], c.store.Employees[found+1:]...)

	return nil
}

func (c *clientImpl) GetEmployees(pagination *coreentity.Pagination) (int, []entity.Employee) {
	var employees []entity.Employee

	var start = (pagination.PageNum - 1) * pagination.PageSize
	var end = pagination.PageNum * pagination.PageSize

	cond := min(end, len(c.store.Employees))
	for i := start; i < cond; i++ {
		employee := c.store.Employees[i]
		employees = append(employees, employee)
	}

	return len(c.store.Employees), employees
}

func (c *clientImpl) GetEmployee(empID int) (*entity.Employee, error) {

	found := c.findEmployeeIndex(empID)
	if found == -1 {
		return nil, fmt.Errorf("employee not found")
	}

	return &c.store.Employees[found], nil
}

func NewClient(store *db.EmployeeStore) Client {
	return &clientImpl{
		store: store,
		index: 1,
	}
}
