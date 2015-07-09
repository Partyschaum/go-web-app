package converters

import (
	"testing"

	"github.com/partyschaum/go-web-app/models"
)

func TestConvertCategoryToViewModel(t *testing.T) {
	category := models.Category{}
	category.SetId(42)
	category.SetImageUrl("the image URL")
	category.SetTitle("the title")
	category.SetDescription("the description")
	category.SetIsOrientRight(true)

	result := ConvertCategoryToViewModel(category)

	if result.Id != category.Id() {
		t.Error("Id not converted properly")
	}

	if result.ImageUrl != category.ImageUrl() {
		t.Error("Image URL not converted properly")
	}

	if result.Title != category.Title() {
		t.Error("Title not converted properly")
	}

	if result.Description != category.Description() {
		t.Error("Description not converted properly")
	}

	if result.IsOrientRight != category.IsOrientRight() {
		t.Error("IsOrientRight not converted properly")
	}
}
