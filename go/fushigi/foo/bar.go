package foo

type bar struct {
	name string
}

func NewBar(name string) *bar {
	return &bar{name}
}

func (b *bar) GetName() string {
	return b.name
}
