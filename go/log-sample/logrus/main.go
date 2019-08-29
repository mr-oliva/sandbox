package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bookun/sandbox/go/log-sample/logrus/subpackage"
	log "github.com/sirupsen/logrus"
)

func init() {
	fmt.Println("init")
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Warn("You should warn")
	subpackage.Sub()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
