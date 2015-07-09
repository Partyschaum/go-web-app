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

type Login struct {
	Base
}

func GetLogin() Login {
	result := Login{
		Base: Base{
			Title: "Lemonade Stand Society - Login",
		},
	}

	return result
}
