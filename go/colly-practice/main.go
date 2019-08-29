package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	detail()
	moreDetail()
}

func index() {
	c := colly.NewCollector()
	c.OnHTML(".newsFeed_item_link", func(ele *colly.HTMLElement) {
		link := ele.Attr("href")
		fmt.Printf("found : %s\n", link)
	})

	c.OnRequest(func(req *colly.Request) {
		fmt.Printf("will request : %s\n", req.URL.String())
	})

	c.Visit("https://news.yahoo.co.jp/topics/top-picks?page=9")
}

func detail() {
	c := colly.NewCollector()
	c.OnHTML(".tpcNews_detailLink > a", func(ele *colly.HTMLElement) {
		link := ele.Attr("href")
		fmt.Printf("found : %s\n", link)
	})
	c.Visit("https://news.yahoo.co.jp/pickup/6334596")
}

func moreDetail() {
	c := colly.NewCollector()
	c.OnHTML(".yjDirectSLinkTarget", func(ele *colly.HTMLElement) {
		contents := ele.DOM.Children().Remove().End().Text()
		fmt.Println(contents)
	})
	c.Visit("https://headlines.yahoo.co.jp/hl?a=20190826-00000020-jij_afp-int")
}
