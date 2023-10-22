package main

import (
	"restful-go/internal/database"
	"restful-go/internal/delivery/rest"
	mRepo "restful-go/internal/repository/menu"
	rUsecase "restful-go/internal/usecase/resto"

	"github.com/labstack/echo/v4"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=root dbname=go_resto_app sslmode=disable"
)

func main() {
	e := echo.New()

	db := database.GetDB(dbAddress)

	menuRepo := mRepo.GetRepository(db)

	restoUsecase := rUsecase.GetUsercase(menuRepo)

	h := rest.NewHandler(restoUsecase)

	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start(":8080"))


}