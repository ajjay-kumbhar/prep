package employee

import (
	coreentity "github.com/ajjay-kumbhar/prep/pkg-core/entity"
	"github.com/ajjay-kumbhar/prep/pkg/service/employee/entity"
)

type Client interface {
	CreateEmployee(employee *entity.Employee) error
	UpdateEmployee(empID int, employee *entity.Employee) error
	DeleteEmployee(empID int) error
	GetEmployees(pagination *coreentity.Pagination) (int, []entity.Employee)
	GetEmployee(empID int) (*entity.Employee, error)
}
