package sample

const (
	OK = iota
	FooError
	BarError
)

func Sample(num int) int {
	if num == 1 {
		return FooError
	}
	if num == 2 {
		return BarError
	}
	return OK
}
