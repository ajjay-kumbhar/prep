package entity

// Employee struct represents an employee.
type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name" validate:"required"`
	Position string  `json:"position" validate:"required"`
	Salary   float64 `json:"salary" validate:"required"`
}
