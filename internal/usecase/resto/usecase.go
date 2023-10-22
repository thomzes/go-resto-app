package resto

import "restful-go/internal/model"

type Usecase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
