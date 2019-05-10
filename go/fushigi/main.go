package main

import (
	"fmt"

	"github.com/bookun/sandbox/go/fushigi/foo"
)

func main() {
	bar := foo.NewBar("fushigi")
	fmt.Println(bar.GetName())
}
