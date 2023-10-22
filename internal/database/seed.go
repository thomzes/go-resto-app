package database

import (
	"restful-go/internal/model"
	"restful-go/internal/model/constant"

	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	foodMenu := []model.MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     30000,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Ayam Rica-rica",
			OrderCode: "ayam_rica_rica",
			Price:     45000,
			Type:      constant.MenuTypeFood,
		},
	}

	drinksMenu := []model.MenuItem{
		{
			Name:      "Es Teh Manis",
			OrderCode: "es_teh_menu",
			Price:     5000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Es Jeruk",
			OrderCode: "es_jeruk",
			Price:     7000,
			Type:      constant.MenuTypeDrink,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(foodMenu)
		db.Create(drinksMenu)
	}
}