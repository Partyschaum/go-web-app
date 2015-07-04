package viewmodels

type Product struct {
	Base
	Name  string
	Price float64
}

func GetProduct() Product {
	return Product{
		Base: Base{
			Title:  "Fesches Produkt",
			Active: "product",
		},
		Name:  "Fesches Produkt",
		Price: 1.99,
	}
}
