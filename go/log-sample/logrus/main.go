package main

import (
	"fmt"
	"os"

	"github.com/bookun/sandbox/go/log-sample/logrus/subpackage"
	log "github.com/sirupsen/logrus"
)

func init() {
	fmt.Println("init")
	//	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {
	log.Warn("You should warn")
	subpackage.Sub()
}
