package converters

import (
	"github.com/partyschaum/go-web-app/models"
	"github.com/partyschaum/go-web-app/viewmodels"
)

func ConvertCategoryToViewModel(category models.Category) viewmodels.Category {
	result := viewmodels.Category{
		Id:            category.Id(),
		ImageUrl:      category.ImageUrl(),
		Title:         category.Title(),
		Description:   category.Description(),
		IsOrientRight: category.IsOrientRight(),
	}
	return result
}
