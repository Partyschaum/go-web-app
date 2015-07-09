package viewmodels

type standLocator struct {
	Base
}

func GetStandLocator() standLocator {
	result := standLocator{
		Base: Base{
			Title:  "Lemonade Stand Supply - Stand Locator",
			Active: "stand_locator",
		},
	}
	return result
}

type standLocation struct {
	Lat   float32
	Lng   float32
	Title string
}

func GetStandLocations() []standLocation {
	result := []standLocation{
		standLocation{
			Lat:   37.41,
			Lng:   -122.102,
			Title: "Lorelei's stand",
		},
		standLocation{
			Lat:   37.412,
			Lng:   -122.099,
			Title: "Rebecca's stand",
		},
		standLocation{
			Lat:   37.407,
			Lng:   -122.1025,
			Title: "Chris's stand",
		},
		standLocation{
			Lat:   37.432,
			Lng:   -122.1025,
			Title: "Carson's stand",
		},
	}
	return result
}
