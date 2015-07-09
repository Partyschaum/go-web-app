package models

type Category struct {
	id            int
	imageUrl      string
	title         string
	description   string
	isOrientRight bool
}

func (c *Category) Id() int {
	return c.id
}

func (c *Category) ImageUrl() string {
	return c.imageUrl
}

func (c *Category) Title() string {
	return c.title
}

func (c *Category) Description() string {
	return c.description
}

func (c *Category) IsOrientRight() bool {
	return c.isOrientRight
}

func (c *Category) SetId(value int) {
	c.id = value
}

func (c *Category) SetImageUrl(value string) {
	c.imageUrl = value
}

func (c *Category) SetTitle(value string) {
	c.title = value
}

func (c *Category) SetDescription(value string) {
	c.description = value
}

func (c *Category) SetIsOrientRight(value bool) {
	c.isOrientRight = value
}

func GetCategories() []Category {
	result := []Category{
		Category{
			id:            1,
			imageUrl:      "lemon.png",
			title:         "Juices and Mixes",
			description:   "Leckerer Limonennektar",
			isOrientRight: false,
		},
		Category{
			id:            2,
			imageUrl:      "kiwi.png",
			title:         "Cups, Straws, and Other Supplies",
			description:   "Wichtiges Zeug",
			isOrientRight: true,
		},
		Category{
			id:            3,
			imageUrl:      "pineapple.png",
			title:         "Signs and Advertising",
			description:   "Werbung and stuff",
			isOrientRight: false,
		},
	}

	return result
}
