package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func serverReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)

	req.URL.Host = "img.news.goo.ne.jp"
	req.URL.Scheme = "http"
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = "img.news.goo.ne.jp"

	proxy.ServeHTTP(res, req)
}

func handler(w http.ResponseWriter, r *http.Request) {
	serverReverseProxy("http://img.news.c0.goo.ne.jp", w, r)
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	//director := func(r *http.Request) {
	//	url := *r.URL
	//	url.Scheme = "https"
	//	url.Host = "img.news.goo.ne.jp"

	//	buf, err := ioutil.ReadAll(r.Body)
	//	if err != nil {
	//		log.Fatal(err.Error())
	//	}
	//	req, err := http.NewRequest(r.Method, url.String(), bytes.NewBuffer(buf))
	//	if err != nil {
	//		log.Fatal(err.Error())
	//	}
	//	req.Header = r.Header
	//	*r = *req
	//	dump, _ := httputil.DumpRequest(req, true)
	//	fmt.Println(dump)
	//}

	//rp := &httputil.ReverseProxy{Director: director}
	//server := http.Server{
	//	Addr:    ":8080",
	//	Handler: rp,
	//}
	//if err := server.ListenAndServe(); err != nil {
	//	log.Fatal(err.Error())
	//}
}
