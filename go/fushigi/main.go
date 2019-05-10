package main

import (
	"fmt"

	"github.com/bookun/sandbox/go/fushigi/foo"
)

func main() {
	bar := foo.NewBar("fushigi")
	bar2 := &foo.Bar{}
	fmt.Println(bar.GetName())
	fmt.Println(bar2.GetName())
}
