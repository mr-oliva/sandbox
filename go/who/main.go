package main

import (
	"log"

	"github.com/domainr/whois"
)

func main() {
	//result, err := whois.Whois("61.126.179.233")
	query := "61.126.179.233"
	request, err := whois.NewRequest(query)
	if err != nil {
		log.Fatal(err)
	}
	response, err := whois.DefaultClient.Fetch(request)
	if err != nil {
		log.Fatal(err)
	}
}
