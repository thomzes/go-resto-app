package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=root dbname=go_resto_app sslmode=disable"
)

type MenuType string

const (
	MenuTypeFood = "food"
	MenuTypeDrink = "drink"
)

type MenuItem struct {
	Name string
	OrderCode string
	Price int
	Type MenuType
}

func getMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)		
	}

	var menuData []MenuItem

	db.Where(MenuItem{Type: MenuType(menuType)}).Find(&menuData)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data" : menuData,
	})
}


func seedDB() {
	foodMenu := []MenuItem{
		{
			Name: "Bakmie",
			OrderCode: "bakmie",
			Price: 30000,
			Type: MenuTypeFood,

		},
		{
			Name: "Ayam Rica-rica",
			OrderCode: "ayam_rica_rica",
			Price: 45000,
			Type: MenuTypeFood,

		},
	}

	drinksMenu := []MenuItem{
		{
			Name: "Es Teh Manis",
			OrderCode: "es_teh_menu",
			Price: 5000,
			Type: MenuTypeDrink,

		},
		{
			Name: "Es Jeruk",
			OrderCode: "es_jeruk",
			Price: 7000,
			Type: MenuTypeDrink,

		},
	}

	fmt.Println(foodMenu, drinksMenu)

	db, err := gorm.Open(postgres.Open(dbAddress))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&MenuItem{})

	if err := db.First(&MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(foodMenu)
		db.Create(drinksMenu)
	}	
}

func main() {
	e := echo.New()

	seedDB()

	e.GET("/menu", getMenu)
	e.Logger.Fatal(e.Start(":8080"))


}