package viewmodels

type Products struct {
	Base
	Products []Product
}

func GetProducts() Products {
	products := Products{
		Base: Base{
			Title:  "Fesches Produkt",
			Active: "product",
		},
	}

	oneProduct := Product{
		Base: Base{
			Title:  "lecker hase",
			Active: "product",
		},
		Name:  "lecker hase",
		Price: 0.99,
	}

	anotherProduct := Product{
		Base: Base{
			Title:  "lecker lolli",
			Active: "product",
		},
		Name:  "lecker lolli",
		Price: 0.22,
	}

	products.Products = []Product{
		oneProduct,
		anotherProduct,
	}

	return products
}
