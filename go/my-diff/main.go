package main

import (
	"fmt"

	"github.com/rsc/diff"
)

func main() {
	//str1 := `sample1
	//hoge123
	//hoge2
	//hoge3
	//`

	str2 := `sample2
    hi
    low`
	str3 := `sample2
    hi
    low`

	if d := diff.Format(str3, str2); d == "" {
		fmt.Println("no diff")
	} else {
		fmt.Println(d)
	}
}
