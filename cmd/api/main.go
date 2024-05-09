package main

import (
	"github.com/labstack/echo/v4"

	"github.com/ajjay-kumbhar/prep/pkg-core/validation"
	"github.com/ajjay-kumbhar/prep/pkg/service/employee/db"

	handler "github.com/ajjay-kumbhar/prep/internal"
	service "github.com/ajjay-kumbhar/prep/pkg/service/employee"
)

func main() {
	e := echo.New()
	e.Validator = validation.NewCustomeValidator()

	// initialize store
	store := db.NewEmployeeStore()
	// initialize client
	client := service.NewClient(store)

	handler.NewHandler(e, client)

	e.Logger.Fatal(e.Start(":8080"))
}
