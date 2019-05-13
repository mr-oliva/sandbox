package sample

const (
	OK = iota
	FooError
	BarError
)

type ErrorCode int

func Sample(num int) ErrorCode {
	if num == 1 {
		return FooError
	}
	if num == 2 {
		return BarError
	}
	return OK
}
