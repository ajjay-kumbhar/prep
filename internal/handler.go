package internal

import (
	"github.com/ajjay-kumbhar/prep/pkg/service/employee"
	"github.com/labstack/echo/v4"
)

func NewHandler(echo *echo.Echo, client employee.Client) {

	eh := newEmployeeHandler(client)

	eg := echo.Group("/v1/employee")
	eg.GET("/", eh.GetEmployees)
	eg.GET("/:empID", eh.GetEmployee)
	eg.POST("/", eh.CreateEmployee)
	eg.PUT("/:empID", eh.UpdateEmployee)
	eg.DELETE("/:empID", eh.DeleteEmployee)
}
