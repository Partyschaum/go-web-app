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
	return categories
}
