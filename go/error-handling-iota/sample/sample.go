package sample

import "fmt"

const (
	OK = iota
	FooError
	BarError
)

type ErrorCode int

func Sample(num int) (ErrorCode, error) {
	if num == 1 {
		return FooError, fmt.Errorf("num=1 is FooError")
	}
	if num == 2 {
		return BarError, fmt.Errorf("num=2 is BarError")
	}
	return OK, nil
}
