package menu

import "restful-go/internal/model"

type Repository interface {
	GetMenu(MenuType string) ([]model.MenuItem, error)
}




