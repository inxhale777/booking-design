package room

type Size string

const (
	Single Size = "single"
	Double Size = "double"
	Triple Size = "triple"
)

var AllSizes = []Size{Single, Double, Triple}
