package viewmodels

const juice int = 1
const supply int = 2
const advertising int = 3

type Products struct {
	Base
	Products []Product
}

func GetProducts(id int) Products {
	var shopName string
	switch id {
	case 1:
		shopName = "Juice"
	case 2:
		shopName = "Supply"
	case 3:
		shopName = "Advertising"
	}

	products := Products{
		Base: Base{
			Title:  "Lemonade Stand Society = " + shopName + " Shop",
			Active: "product",
		},
	}

	if id == juice {
		products.Products = getProductList()
	}

	return products
}

func getProductList() []Product {
	lemonJuice := Product{
		Base: Base{
			Title:  "Lemon Juice",
			Active: "product",
		},
		Name:  "Lemon Juice",
		Price: 0.99,
	}

	appleJuice := Product{
		Base: Base{
			Title:  "Apple Juice",
			Active: "product",
		},
		Name:  "Apple Juice",
		Price: 0.22,
	}

	watermelonJuice := Product{
		Base: Base{
			Title:  "Watermelon Juice",
			Active: "product",
		},
		Name:  "Watermelon Juice",
		Price: 0.22,
	}

	kiwiJuice := Product{
		Base: Base{
			Title:  "Kiwi Juice",
			Active: "product",
		},
		Name:  "Kiwi Juice",
		Price: 0.22,
	}

	mangosteenJuice := Product{
		Base: Base{
			Title:  "Mangosteen Juice",
			Active: "product",
		},
		Name:  "Mangosteen Juice",
		Price: 0.22,
	}

	orangeJuice := Product{
		Base: Base{
			Title:  "Orange Juice",
			Active: "product",
		},
		Name:  "Orange Juice",
		Price: 0.22,
	}

	pineappleJuice := Product{
		Base: Base{
			Title:  "Pineapple Juice",
			Active: "product",
		},
		Name:  "Pineapple Juice",
		Price: 0.22,
	}

	strawberryJuice := Product{
		Base: Base{
			Title:  "Strawberry Juice",
			Active: "product",
		},
		Name:  "Strawberry Juice",
		Price: 0.22,
	}

	result := []Product{
		lemonJuice,
		appleJuice,
		watermelonJuice,
		kiwiJuice,
		mangosteenJuice,
		orangeJuice,
		pineappleJuice,
		strawberryJuice,
	}

	return result
}
