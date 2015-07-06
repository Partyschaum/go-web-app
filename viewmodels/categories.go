package viewmodels

type Categories struct {
	Base
	Categories []Category
}

type Category struct {
	Id            int
	ImageUrl      string
	Title         string
	Description   string
	IsOrientRight bool
}

func GetCategories() Categories {
	categories := Categories{
		Base: Base{
			Title:  "Lemonade Stand Society - Shop",
			Active: "shop",
		},
	}

	juiceCategory := Category{
		Id:            1,
		ImageUrl:      "lemon.png",
		Title:         "Juices and Mixes",
		Description:   "Leckerer Limonennektar",
		IsOrientRight: false,
	}

	supplyCategory := Category{
		Id:            2,
		ImageUrl:      "kiwi.png",
		Title:         "Cups, Straws, and Other Supplies",
		Description:   "Wichtiges Zeug",
		IsOrientRight: true,
	}

	advertiseCategory := Category{
		Id:            3,
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
