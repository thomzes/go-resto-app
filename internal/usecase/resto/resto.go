package resto

import (
	"restful-go/internal/model"
	"restful-go/internal/repository/menu"
)

type restoUsecase struct {
	menuRepo menu.Repository
}

func GetUsercase(menuRepo menu.Repository) Usecase {
	return &restoUsecase{
		menuRepo: menuRepo,
	}
}

func (r *restoUsecase) GetMenu(menuType string) ([]model.MenuItem,error) {
	return r.menuRepo.GetMenu(menuType)
}


