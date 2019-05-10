package foo

type Bar struct {
	name string
}

// NewBar return Bar
func NewBar(name string) *Bar {
	return &Bar{name}
}

func (b *Bar) GetName() string {
	return b.name
}
