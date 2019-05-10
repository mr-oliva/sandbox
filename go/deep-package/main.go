package main

import (
	"fmt"

	"github.com/bookun/sandbox/go/deep-package/driver/database"
)

func main() {
	mysql := &database.MySQL{}
	fmt.Println(mysql)
}
