package viewmodels

type Home struct {
	Title  string
	Active string
}

func GetHome() Home {
	return Home{
		Title:  "Lemonade Stand Society",
		Active: "home",
	}
}
