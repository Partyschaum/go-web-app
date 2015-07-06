package viewmodels

type Product struct {
	Base
	Id    int
	Name  string
	Price float64
}

func GetProduct(id int) Product {
	productList := getProductList()

	var product Product
	for _, p := range productList {
		if p.Id == id {
			product = p
			break
		}
	}
	product.Base = Base{
		Title:  "Lemonade Stand Society - " + product.Name,
		Active: "product",
	}

	return product
}
