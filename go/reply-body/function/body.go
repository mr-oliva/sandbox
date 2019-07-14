package function

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func GetBody(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	log.Println(buf.String())
	fmt.Fprint(w, buf.String())
}
