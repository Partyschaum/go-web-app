package viewmodels

type Categories struct {
	Title      string
	Active     string
	Categories []Category
}

type Category struct {
	ImageUrl      string
	Title         string
	Description   string
	IsOrientRight bool
}

func GetCategories() Categories {
	categories := Categories{
		Title:  "Lemonade Stand Society - Shop",
		Active: "shop",
	}

	juiceCategory := Category{
		ImageUrl:      "lemon.png",
		Title:         "Juices and Mixes",
		Description:   "Leckerer Limonennektar",
		IsOrientRight: false,
	}

	supplyCategory := Category{
		ImageUrl:      "kiwi.png",
		Title:         "Cups, Straws, and Other Supplies",
		Description:   "Wichtiges Zeug",
		IsOrientRight: true,
	}

	advertiseCategory := Category{
		ImageUrl:      "pineapple.png",
		Title:         "Signs and Advertising",
		Description:   "Werbung and stuff",
		IsOrientRight: false,
	}

	categories.Categories = []Category{
		juiceCategory,
		supplyCategory,
		advertiseCategory,
	}

	return categories
}
