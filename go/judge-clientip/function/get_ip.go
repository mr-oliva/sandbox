package function

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/bookun/sandbox/go/judge-clientip/driver/cache"
	"github.com/bookun/sandbox/go/judge-clientip/entity"
)

var (
	regMap = map[string]*regexp.Regexp{
		"Flet's hikari":  regexp.MustCompile(os.Getenv("fletsRe")),
		"ipoe":           regexp.MustCompile(os.Getenv("ipoeRe")),
		"B flet's":       regexp.MustCompile(os.Getenv("bFletsRe")),
		"Flet's Premium": regexp.MustCompile(os.Getenv("fletsPremiumRe")),
		"Mobile ONE":     regexp.MustCompile(os.Getenv("mobileOne")),
	}
)

type cacher interface {
	Get(context.Context, string) (entity.Result, error)
	Add(context.Context, string, entity.Result) error
}

type hostResearch struct {
	cache cacher
}

func GetIP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	service := r.URL.Query().Get("service")
	if service == "" {
		log.Println("not found: service parameter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var resultBuf bytes.Buffer

	clientIP := r.Header.Get("X-Forwarded-For")
	if clientIP == "" {
		clientIP = "222.145.13.136"
	}
	clientIP = strings.Split(clientIP, ",")[0]

	ctx := context.Background()

	//cache, err := cache.NewFirebase(ctx, os.Getenv("projectID"))
	//if err != nil {
	//	log.Println(err.Error())
	//	return
	//}
	//cache := cache.NewMemcache(os.Getenv("memcacheServer"))
	cache, err := cache.NewRedis(os.Getenv("redisAddr"), os.Getenv("redisPassword"))
	if err != nil {
		log.Println(err.Error())
		return
	}
	hostResearch := &hostResearch{cache: cache}
	cacheResult, err := hostResearch.cache.Get(ctx, clientIP)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if !cacheResult.IsEmpty() {
		if err := json.NewEncoder(&resultBuf).Encode(cacheResult); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("ip: %s\thost: %s\tkind: %s\tservice: %s\n", cacheResult.IP, cacheResult.Host, cacheResult.Kind, service)
		fmt.Fprint(w, resultBuf.String())
		return
	}

	googleDNSDialer := func(ctx context.Context, network, address string) (net.Conn, error) {
		d := net.Dialer{}
		return d.DialContext(ctx, "udp", "8.8.8.8:53")
	}
	resolver := net.Resolver{PreferGo: true, Dial: googleDNSDialer}
	hosts, err := resolver.LookupAddr(ctx, clientIP)
	if err != nil {
		result := entity.Result{IP: clientIP, Host: "-", Kind: "no_set", Error: err.Error(), Service: service}
		if err := json.NewEncoder(&resultBuf).Encode(result); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, resultBuf.String())
		log.Printf("ip: %s\thost: %s\tkind: %s\tservice: %s\n", clientIP, "-", "no_set", service)
		return
	}

	host := strings.Join(hosts, ",")
	kind := getKind(host)

	result := entity.Result{IP: clientIP, Host: host, Kind: kind, Error: "", Service: service}
	if err := hostResearch.cache.Add(ctx, clientIP, result); err != nil {
		log.Println(err.Error())
	}
	result.From = "dns"
	if err := json.NewEncoder(&resultBuf).Encode(result); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("ip: %s\thost: %s\tkind: %s\tservice: %s\n", clientIP, host, kind, service)
	fmt.Fprint(w, resultBuf.String())
}

func getKind(host string) string {
	for provider, hostPattern := range regMap {
		if hostPattern.MatchString(host) {
			return provider
		}
	}
	return "other"
}
