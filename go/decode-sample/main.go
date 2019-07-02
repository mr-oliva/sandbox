package main

import (
	"fmt"
	"image"
	"log"
	"net/http"

	_ "image/png"
)

func main() {
	sampleURL := "https://3.bp.blogspot.com/-nWrdwPzMYxA/XJiJLA-xQ_I/AAAAAAABSFU/O2yycOXkDeQYF9-dsQqfJi73mhsYUm70wCLcBGAs/s180-c/nigaoe_hattori_hanzou.png"
	resp, err := http.Get(sampleURL)
	if err != nil {
		log.Fatal(err)
	}

	//buf := new(bytes.Buffer)
	//io.Copy(buf, resp.Body)

	_, encode, err := image.DecodeConfig(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(encode)
}
