package internal

import (
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"

	coreentity "github.com/ajjay-kumbhar/prep/pkg-core/entity"

	"github.com/ajjay-kumbhar/prep/pkg/service/employee"
	"github.com/ajjay-kumbhar/prep/pkg/service/employee/entity"
)

type handler struct {
	client employee.Client
}

func newEmployeeHandler(client employee.Client) *handler {
	return &handler{
		client: client,
	}
}

func (h *handler) CreateEmployee(c echo.Context) error {
	dto := &entity.Employee{}
	if err := c.Bind(dto); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "failed to bind employee"}
	}

	if err := c.Validate(dto); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	if err := h.client.CreateEmployee(dto); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"data": "employee created successfully",
	})
}

func (h *handler) UpdateEmployee(c echo.Context) error {
	empIDStr := c.Param("empID")
	if empIDStr == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid employee ID"}
	}

	dto := &entity.Employee{}
	if err := c.Bind(dto); err != nil {
		return err
	}

	if err := c.Validate(dto); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	empID, err := strconv.Atoi(empIDStr)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	if err := h.client.UpdateEmployee(empID, dto); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "employee updated successfully",
	})
}

func (h *handler) DeleteEmployee(c echo.Context) error {
	empIDStr := c.Param("empID")
	if empIDStr == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid employee ID"}
	}

	empID, err := strconv.Atoi(empIDStr)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	if err := h.client.DeleteEmployee(empID); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"data": "employee deleted successfully",
	})
}

func (h *handler) GetEmployees(c echo.Context) error {
	var pagination coreentity.Pagination
	if err := c.Bind(&pagination); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}
	pagination.Validate()

	total, employees := h.client.GetEmployees(&pagination)

	pagination.Total = total
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"data":       employees,
		"pagination": pagination,
	})
}

func (h *handler) GetEmployee(c echo.Context) error {
	empIDStr := c.Param("empID")
	if empIDStr == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid employee ID"}
	}

	empID, err := strconv.Atoi(empIDStr)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	var employee *entity.Employee
	if employee, err = h.client.GetEmployee(empID); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"data": employee,
	})
}
