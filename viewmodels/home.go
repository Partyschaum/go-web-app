package viewmodels

type Home struct {
	Base
}

func GetHome() Home {
	return Home{
		Base: Base{
			Title:  "Lemonade Stand Society",
			Active: "home",
		},
	}
}
