package room

type Category string

const (
	Economy Category = "economy"
	Comfort Category = "comfort"
	Lux     Category = "lux"
)

var AllCategories = []Category{Economy, Comfort, Lux}
